#! /usr/bin/python3
import json
import pdb
import iota.harness.api as api

def ServerCmd(port = None, time=None, run_core=None):
    assert(port)
    cmd = ["iperf3", "-s","-p", str(port)]
    if run_core:
        cmd.extend(["-A", str(run_core)])

    if time:
        cmd.extend(["-t", str(time)])

    return " ".join(cmd)

def ClientCmd(server_ip, port = None, time=10, pktsize=None, proto='tcp', run_core=None,
            ipproto='v4', num_of_streams = None, jsonOut=False, connect_timeout=None):
    assert(port)
    cmd = ["iperf3", "-c", str(server_ip), "-p", str(port)]

    if time:
        cmd.extend(["-t", str(time)])

    if run_core:
        cmd.extend(["-A", str(run_core)])

    if proto == 'udp':
        cmd.append('-u')

    if jsonOut:
        cmd.append('-J')

    if num_of_streams:
        cmd.extend(["-P", str(num_of_streams)])

    if connect_timeout:
        cmd.extend(["--connect-timeout", str(connect_timeout)])

    if pktsize:
        if proto == 'tcp':
            cmd.extend(["-M", str(pktsize)])
        else:
            cmd.extend(["-l", str(pktsize)])

    if ipproto == 'v6':
        cmd.append("-6")

    return " ".join(cmd)


def __get_json(iperf_out):
    try:
        return json.loads(iperf_out)
    except:
        api.Logger.error("Failed to parse iperf json output :", iperf_out)
        return {"error" : "Json Parse error"}

def ConnectionRefused(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "Connection refused" in error:
        return True
    return False

def ConnectionTimedout(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "Connection timed out" in error:
        return True
    return False

def ControlSocketClosed(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "control socket has closed unexpectedly" in error:
        return True
    return False

def UnknownControlMessage(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "unknown control message" in error:
        return True
    return False


def ConnectionSuccessful(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if not error:
        return True
    api.Logger.error("Connection Error :",  error)
    return False

def ServerTerminated(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "the server has terminated" in error:
        return True
    return False

def ClientTerminated(iperf_out):
    iperfJson = __get_json(iperf_out)
    error = iperfJson.get("error", None)
    if error and "the client has terminated" in error:
        return True
    return False

def Success(iperf_out):
    iperfJson = __get_json(iperf_out)
    return iperfJson.get("error", None) is None

def Error(iperf_out):
    iperfJson = __get_json(iperf_out)
    return iperfJson.get("error", None)

def GetSentGbps(iperf_out):
    iperfJson = __get_json(iperf_out)
    end = iperfJson.get("end", None)
    if not end:
        return '0'
    if iperfJson["start"]["test_start"]["protocol"] == 'TCP':
        return ('{0:.2f}'.format(end["sum_sent"]["bits_per_second"]/(1024 * 1024 * 1024)))
    return ('{0:.2f}'.format(end["sum"]["bits_per_second"]/(1024 * 1024 * 1024)))


def GetReceivedGbps(iperf_out):
    iperfJson = __get_json(iperf_out)
    end = iperfJson.get("end", None)
    if not end:
        return '0'
    if iperfJson["start"]["test_start"]["protocol"] == 'TCP':
        return ('{0:.2f}'.format(end["sum_received"]["bits_per_second"]/(1024 * 1024 * 1024)))
    return '0'


if __name__ == '__main__':
    out = open("/Users/sudhiaithal/Downloads/test.json", "r")
    iperf_out = out.read()
    if Success(iperf_out):
        print ("Server terminated")
    print (GetSentGbps(iperf_out))
    print (GetReceivedGbps(iperf_out))
