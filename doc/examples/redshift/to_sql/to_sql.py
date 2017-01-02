#!/usr/bin/python3

import sys, json

# Sanity check arguments -- need name of database table
if len(sys.argv) != 2:
    sys.stderr.write("Error: to_sql.py takes 1 argument: the name of the "
            +"table to write values into\n")
    sys.exit(1)

# Parse stdin as a stream of 1-layer deep json objects
def next_obj(stream):
    result = []
    for l in stream.readlines():
        idx = l.find("}")
        if idx >= 0:
            result.append(l[:idx+1])
            yield json.loads("".join(result))
            result = [l[idx+1:]]
        else:
            result.append(l)

for obj in next_obj(sys.stdin):
    print("INSERT INTO {} ({}) VALUES ({})".format(
        sys.argv[1],
        ", ".join(obj.keys()),
        ", ".join([ str(v) for v in obj.values() ])))
