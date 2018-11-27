#! /usr/bin/python3
import pdb
import traceback

import iota.harness.infra.store as store
import iota.harness.infra.types as types
import iota.harness.infra.utils.parser as parser
import iota.harness.infra.utils.loader as loader
import iota.harness.api as api
import iota.harness.infra.testcase as testcase
import iota.harness.infra.topology as topology
import iota.harness.infra.logcollector as logcollector
import iota.harness.infra.utils.utils as utils
import iota.harness.infra.utils.timeprofiler as timeprofiler

from iota.harness.infra.utils.logger import Logger as Logger
from iota.harness.infra.glopts import GlobalOptions as GlobalOptions

class TestSuite:
    def __init__(self, spec):
        self.__spec = spec
        self.__tcs = []

        if GlobalOptions.testsuites and self.Name() not in GlobalOptions.testsuites:
            Logger.debug("Skipping Testsuite: %s because of command-line filters." % self.Name())
            self.__enabled = False
            return

        #if GlobalOptions.mode != self.Mode():
        #    Logger.info("Skipping Testsuite: %s because of mode: %s" % (self.Name(), self.Mode()))
        #    self.__enabled = False
        #    return

        self.__enabled = getattr(self.__spec.meta, 'enable', True)
        self.__aborted = False
        self.__attrs = {}
        self.__timer = timeprofiler.TimeProfiler()

        self.__stats_pass = 0
        self.__stats_fail = 0
        self.__stats_ignored = 0
        self.__stats_error = 0
        self.__stats_total = 0
        self.result = types.status.FAILURE
        return

    def Abort(self):
        self.__aborted = True
        self.__curr_tc.Abort()
        return

    def GetTestbedType(self):
        if self.__spec.meta.mode.lower() == 'hardware':
            return types.tbtype.HARDWARE
        elif self.__spec.meta.mode.lower() == 'simulation':
            return types.tbtype.SIMULATION
        elif self.__spec.meta.mode.lower() == 'hybrid':
            return types.tbtype.HYBRID
        return types.tbtype.ANY

    def GetImages(self):
        return self.__spec.images

    def GetTopology(self):
        return self.__topology

    def GetNicMode(self):
        return self.__spec.meta.nicmode

    def IsConfigOnly(self):
       return getattr(self.__spec.meta, "cfgonly", False)
 
    def DoConfig(self):
        return self.__setup_config()

    def __import_testbundles(self):
        if getattr(self.__spec, 'testcases', None) == None:
            self.__spec.testcases = []

        for imp in self.__spec.testbundles:
            filename = "%s/iota/test/iris/testbundles/%s" % (api.GetTopDir(), imp)
            tcb = parser.YmlParse(filename)
            tcb.meta.os = getattr(tcb.meta, 'os', [ 'linux' ])
            if store.GetTestbed().GetOs() not in tcb.meta.os:
                Logger.debug("Skipping Testbundle: %s due to OS mismatch." % tcb.meta.name)
                continue
            if GlobalOptions.testbundles and tcb.meta.name not in  GlobalOptions.testbundles:
                Logger.info("Skipping Testbundle: %s due to cmdline filter." % tcb.meta.name)
                continue
            self.__spec.testcases.extend(tcb.testcases)
            for tc in tcb.testcases:
                tc.bundle = tcb.meta.name
        return

    def __resolve_testcases(self):
        for tc_spec in self.__spec.testcases:
            if getattr(tc_spec, 'disable', False):
                Logger.info("Skipping disabled test case %s" % tc_spec.name)
                continue
            tc_spec.packages = self.__spec.packages
            if getattr(self.__spec, 'common', None) and getattr(self.__spec.common, 'verifs', None):
                tc_spec.verifs = self.__spec.common.verifs
            tc = testcase.Testcase(tc_spec)
            self.__tcs.append(tc)
        return types.status.SUCCESS

    def __resolve_teardown(self):
        teardown_spec = getattr(self.__spec, 'teardown', [])
        if teardown_spec is None:
            return types.status.SUCCESS
        for s in self.__spec.teardown:
            Logger.debug("Resolving teardown module: %s" % s.step)
            s.step = loader.Import(s.step, self.__spec.packages)
        return types.status.SUCCESS

    def __expand_iterators(self):
        return

    def __parse_setup_topology(self):
        topospec = getattr(self.__spec.setup, 'topology', None)
        if not topospec:
            Logger.error("Error: No topology specified in the testsuite.")
            assert(0)
        self.__topology = topology.Topology(topospec)
        return types.status.SUCCESS

    def __resolve_setup_config(self):
        cfgspec = getattr(self.__spec.setup, 'config', None)
        if not cfgspec:
            return types.status.SUCCESS
        for s in self.__spec.setup.config:
            Logger.debug("Resolving config step: %s" % s.step)
            s.step = loader.Import(s.step, self.__spec.packages)
        return types.status.SUCCESS

    def __parse_setup(self):
        ret = self.__parse_setup_topology()
        if ret != types.status.SUCCESS:
            return ret

        ret = self.__resolve_setup_config()
        if ret != types.status.SUCCESS:
            return ret
        return types.status.SUCCESS

    def __setup_config(self):
        for s in self.__spec.setup.config:
            status = loader.RunCallback(s.step, 'Main', True, None)
            if status != types.status.SUCCESS:
                return status
        return types.status.SUCCESS

    def __setup(self):
        ret = self.__topology.Setup(self)
        if ret != types.status.SUCCESS:
            return ret
        ret = self.__setup_config()
        if ret != types.status.SUCCESS:
            return ret

        return types.status.SUCCESS

    def __update_stats(self):
        for tc in self.__tcs:
            p,f,i,e = tc.GetStats()
            self.__stats_pass += p
            self.__stats_fail += f
            self.__stats_ignored += i
            self.__stats_error += e
            
        self.__stats_total = (self.__stats_pass + self.__stats_fail +\
                              self.__stats_ignored + self.__stats_error)
        return

    def __execute_testcases(self):
        result = types.status.SUCCESS
        for tc in self.__tcs:
            self.__curr_tc = tc
            ret = tc.Main()
            if ret != types.status.SUCCESS:
                result = ret
                if GlobalOptions.no_keep_going:
                    return ret
            if self.__aborted:
                return types.status.FAILURE
        return result

    def Name(self):
        return self.__spec.meta.name

    def Mode(self):
        return self.__spec.meta.mode

    def LogsDir(self):
        return "%s/iota/logs/%s" % (api.GetTopDir(), self.Name())

    def __get_oss(self):
        return getattr(self.__spec.meta, 'os', ['linux'])

    def Main(self):
        if not self.__enabled:
           return types.status.SUCCESS

        if self.GetTestbedType() != types.tbtype.ANY and\
           self.GetTestbedType() != store.GetTestbed().GetTestbedType():
           Logger.info("Skipping Testsuite: %s due to testbed type mismatch." % self.Name())
           self.__enabled = False
           return types.status.SUCCESS

        if store.GetTestbed().GetOs() not in self.__get_oss():
            Logger.info("Skipping Testsuite: %s due to OS mismatch." % self.Name())
            self.__enabled = False
            return types.status.SUCCESS
        
        # Start the testsuite timer
        self.__timer.Start()
        
        # Update logger
        Logger.SetTestsuite(self.Name())
        Logger.info("Starting Testsuite: %s" % self.Name())

        # Initialize Testbed for this testsuite
        status = store.GetTestbed().InitForTestsuite(self)
        if status != types.status.SUCCESS:
            self.__timer.Stop()
            return status
        
        # Use try/except block to not let testcase crash the overall run.
        self.__import_testbundles()
        self.__resolve_testcases()
        self.__resolve_teardown()
        self.__expand_iterators()

        status = self.__parse_setup()
        if status != types.status.SUCCESS:
            self.__timer.Stop()
            return status

        status = self.__setup()
        if status != types.status.SUCCESS:
            self.__timer.Stop()
            return status
    
        try:
            self.result = self.__execute_testcases()
        except:
            utils.LogException(Logger)
            Logger.error("EXCEPTION: Aborting Testcase Execution.")
            self.result = types.status.FAILURE

        self.__update_stats()
        Logger.info("Testsuite %s FINAL STATUS = %d" % (self.Name(), self.result))
        
        logcollector.CollectLogs()
        self.__timer.Stop()
        return self.result

    def PrintReport(self):
        if not self.__enabled:
           return types.status.SUCCESS
        print("\nTestSuite: %s" % self.__spec.meta.name)
        print(types.HEADER_SUMMARY)
        print(types.FORMAT_TESTCASE_SUMMARY %\
              ("Testbundle", "Testcase", "Owner", "Result", "Duration"))
        print(types.HEADER_SUMMARY)
        for tc in self.__tcs:
            tc.PrintResultSummary()
        print(types.HEADER_SUMMARY)

    def PrintSummary(self):
        if not self.__enabled:
           return types.status.SUCCESS
        print(types.FORMAT_ALL_TESTSUITE_SUMMARY %\
              (self.__spec.meta.name, self.__stats_pass, self.__stats_fail, self.__stats_ignored,
               self.__stats_error, self.__stats_total, types.status.str(self.result).title(),
               self.__timer.TotalTime()))
        return types.status.SUCCESS

    def SetAttr(self, attr, value):
        Logger.info("Adding new Testsuite attribute: %s = " % attr, value)
        self.__attrs[attr] = value
        return

    def GetAttr(self, attr):
        return self.__attrs[attr]
