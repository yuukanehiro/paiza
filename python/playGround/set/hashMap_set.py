import random
import time

# Set lookup time: 2.375 seconds
huge_list1 = list(range(1, 100000000))
query_list = [random.randint(1, 100000000) for _ in range(50)]
start_time = time.time()
huge_list_set1 = set(huge_list1)
for q in query_list:
    if q in huge_list_set1:
        print(f"YES! q:{q}")
end_time = time.time()
print(f"Set lookup time: {end_time - start_time:.3f} seconds")


# List lookup time: 20.923 seconds
huge_list2 = list(range(1, 100000000))
query_list2 = [random.randint(1, 100000000) for _ in range(50)]

start_time = time.time()
for q in query_list2:
    if q in huge_list2:
        print(f"YES! q:{q}")
end_time = time.time()
print(f"List lookup time: {end_time - start_time:.3f} seconds")

# % python3 hashMap_set.py
# YES! q:6425418
# YES! q:55254154
# YES! q:79766488
# YES! q:61209355
# YES! q:39598799
# YES! q:16548314
# YES! q:37099041
# YES! q:11282806
# YES! q:41110452
# YES! q:98703472
# YES! q:71902607
# YES! q:29507969
# YES! q:94818108
# YES! q:84847321
# YES! q:98153957
# YES! q:99022483
# YES! q:91775775
# YES! q:933102
# YES! q:38677511
# YES! q:33840010
# YES! q:32967669
# YES! q:29688017
# YES! q:4873671
# YES! q:6162259
# YES! q:30750605
# YES! q:26906363
# YES! q:94799914
# YES! q:51329903
# YES! q:76654843
# YES! q:29778757
# YES! q:64300180
# YES! q:1662239
# YES! q:25415697
# YES! q:51386454
# YES! q:43049078
# YES! q:34147660
# YES! q:40097173
# YES! q:39638912
# YES! q:67817757
# YES! q:28000363
# YES! q:83706911
# YES! q:87045306
# YES! q:54727864
# YES! q:33360289
# YES! q:41517357
# YES! q:65972729
# YES! q:72994186
# YES! q:36452246
# YES! q:57005056
# YES! q:43336860
# Set lookup time: 2.375 seconds


# YES! q:93829146
# YES! q:95608236
# YES! q:769081
# YES! q:45108706
# YES! q:15183555
# YES! q:17827046
# YES! q:88264608
# YES! q:20788918
# YES! q:82474918
# YES! q:80027019
# YES! q:6904438
# YES! q:15283659
# YES! q:91762097
# YES! q:33487946
# YES! q:41232917
# YES! q:46809925
# YES! q:62912945
# YES! q:22085457
# YES! q:22363228
# YES! q:28880254
# YES! q:72801225
# YES! q:29735913
# YES! q:22869082
# YES! q:80918932
# YES! q:61914846
# YES! q:15753799
# YES! q:83076284
# YES! q:70389067
# YES! q:99326400
# YES! q:1274609
# YES! q:30649847
# YES! q:15648691
# YES! q:10300567
# YES! q:58403214
# YES! q:43215400
# YES! q:66164048
# YES! q:37326270
# YES! q:67963996
# YES! q:36293965
# YES! q:28247061
# YES! q:97120432
# YES! q:71236631
# YES! q:90308545
# YES! q:50140941
# YES! q:30221195
# YES! q:96139561
# YES! q:14214527
# YES! q:18888348
# YES! q:64382815
# YES! q:82068594
# List lookup time: 20.923 seconds