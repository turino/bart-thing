import sys

f = open("args.txt", "w")
f.write(", ".join(sys.argv))
f.close()
