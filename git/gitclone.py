#!/usr/bin/env python
# coding=utf-8 

# for git clone helper
# author: jason fang
# date: 20160624

import os, sys

GERRIT_SERVER="192.168.1.187"
GERRIT_LISTEN_PORT="29419"
GIT_REPO_NAME=""
SSH_CLONE_FULL_CMD=""
GERRIT_HOOKS_CMD="" 
WHOAMI=""

def hello():
	print '\033[0;32;40m'
	#print 50*('*')
	print 
	print '\tGerrit source code clone tools'
	print 
	print 50*("*")
	print '\033[0m'

def __exit(x):
	sys.exit(x)
	
def error_color():
	print '\033[0;31;40m'
	
def normal_color():
	print '\033[0m'
	
def get_username():
	global WHOAMI
	WHOAMI = os.popen('whoami').read().strip()
	
	if cmp(WHOAMI,"jason") == 0:
		WHOAMI="fangjianjiang"
	
	
def prep_git_repo():
	global GIT_REPO_NAME
	GIT_REPO_NAME=raw_input("Please input the dest Git Repository:")
	if (len(GIT_REPO_NAME) <= 0):
		error_color()
		print "Invalid Git Repository, You entered NULL value!"
		normal_color()
		__exit(1)
	
	print("The dest Git Repo is %s") % GIT_REPO_NAME

	
def prep_ssh_clone_str():
	global SSH_CLONE_FULL_CMD
	global GIT_REPO_NAME
	global WHOAMI 
	global GERRIT_SERVER
	global GERRIT_LISTEN_PORT
	SSH_CLONE_FULL_CMD = "git clone ssh://"+WHOAMI+"@"+GERRIT_SERVER+":"+GERRIT_LISTEN_PORT+"/"+GIT_REPO_NAME
	print SSH_CLONE_FULL_CMD
	
def prep_gerrit_hooks():
	global WHOAMI
	global GIT_REPO_NAME
	global GERRIT_HOOKS_CMD
	global GERRIT_LISTEN_PORT
	global GERRIT_SERVER
	GERRIT_HOOKS_CMD = "scp -p -P " + GERRIT_LISTEN_PORT + " " + WHOAMI+ "@" + GERRIT_SERVER + ":hooks/commit-msg " + GIT_REPO_NAME + "/.git/hooks/"
	print GERRIT_HOOKS_CMD
	
def do_shell_cmd(CMD):
	os.system(CMD)
	
def do_git_clone():	
	global GERRIT_HOOKS_CMD
	global SSH_CLONE_FULL_CMD
	
	get_username()
	prep_git_repo()
	prep_ssh_clone_str()
	prep_gerrit_hooks()
	
	do_shell_cmd(SSH_CLONE_FULL_CMD)
	do_shell_cmd(GERRIT_HOOKS_CMD)
	
	print '\033[0;32;40m'
	print "Git Clone done!"
	normal_color()
	
if __name__ == "__main__":
	hello()
	do_git_clone()
	__exit(0)





