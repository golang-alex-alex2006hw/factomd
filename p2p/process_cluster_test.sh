#!/bin/bash
echo "run this from factomproject/factomd eg:"
echo "$ ./p2p/process_cluster_test.sh"
echo
echo "Compiling..."
go install -a 
if [ $? -eq 0 ]; then
    echo "Running..."
    factomd -count=1 -folder="test1-" -port 9120 -p2pAddress="tcp://127.0.0.1:34340" -peers="tcp://127.0.0.1:34341 tcp://127.0.0.1:34342 tcp://127.0.0.1:34343" -leader=true & node0=$!
    factomd -count=1 -folder="test2-" -port 9121 -p2pAddress="tcp://127.0.0.1:34341" -peers="tcp://127.0.0.1:34340 tcp://127.0.0.1:34342 tcp://127.0.0.1:34343" -follower=true & node1=$!
#    factomd -count=1 -folder="test3-" -port 9122 -p2pAddress="tcp://127.0.0.1:34342" -peers="tcp://127.0.0.1:34340 tcp://127.0.0.1:34341 tcp://127.0.0.1:34343" -follower=true & node2=$!
#    factomd -count=1 -folder="test4-" -port 9123 -p2pAddress="tcp://127.0.0.1:34343" -peers="tcp://127.0.0.1:34340 tcp://127.0.0.1:34341 tcp://127.0.0.1:34342" -follower=true & node3=$!
    sleep 60
    kill $node0 $node1 $node2 $node3
fi