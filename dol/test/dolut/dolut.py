#! /usr/bin/python3

def Setup(infra, module):
    iterelem = module.iterator.Get()

    if iterelem:
        if 'root' in iterelem.__dict__:
            #Alterative would be to do Get on root by passing ConfigStore
            #However this is testspec, not a good idea to import.
            module.testspec.selectors.root.re_init(iterelem.root)

        if 'tenant' in iterelem.__dict__:
            module.testspec.selectors.tenant.Extend(iterelem.tenant)

        if 'network' in iterelem.__dict__:
            module.testspec.selectors.network.Extend(iterelem.network)

        if 'segment' in iterelem.__dict__:
            module.testspec.selectors.segment.Extend(iterelem.segment)

        if 'endpoint' in iterelem.__dict__:
            module.testspec.selectors.endpoint.Extend(iterelem.endpoint)
            
        if 'enic' in iterelem.__dict__:
            module.testspec.selectors.enic.Extend(iterelem.enic)

        if 'lif' in iterelem.__dict__:
            module.testspec.selectors.lif.Extend(iterelem.lif)

        if 'uplink' in iterelem.__dict__:
            module.testspec.selectors.uplink.Extend(iterelem.uplink)

        if 'uplinkpc' in iterelem.__dict__:
            module.testspec.selectors.uplinkpc.Extend(iterelem.uplinkpc)

        if 'security_profile' in iterelem.__dict__:
            module.testspec.selectors.security_profile.Extend(iterelem.security_profile)

        if 'session' in iterelem.__dict__:
            module.testspec.selectors.session.base.Extend(iterelem.session.base)
            module.testspec.selectors.session.rflow.Extend(iterelem.session.rflow)
            module.testspec.selectors.session.iflow.Extend(iterelem.session.iflow)

    if module.args == None:
        return

    if 'maxflows' in module.args.__dict__:
        module.testspec.selectors.maxflows = module.args.maxflows

    return

def Teardown(infra, module):
    print("Teardown(): Sample Implementation.")
    return

def TestCaseSetup(tc):
    print("TestCaseSetup(): Sample Implementation.")
    return

def TestCaseVerify(tc):
    print("TestCaseVerify(): Sample Implementation.")
    return True

def TestCaseTeardown(tc):
    print("TestCaseTeardown(): Sample Implementation.")
    return

