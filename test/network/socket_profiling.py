import os
import subprocess

def main(): 
 #   start_cedana_daemon()
    # daemon is running, tell it to start process 
    pid = 70992
    

def start_cedana_daemon():
    subprocess.Popen(["sudo", "./../../cedana", "daemon"], shell=False)

def exec_cedana_process(process): 
    result = subprocess.Popen(["sudo", "./../../cedana", "start", "-p", process], shell=False)
    print(result.pid)
main()