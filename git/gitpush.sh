#!/bin/bash
# Git push gerrit tools
# version: 1.0


ACTION=push
PARM=origin
GIT_REFERENCE_PATH="HEAD:refs/for"
NAME=
REVIEWS=
W=0
M=master

export PATH=$PATH

welcome(){
	echo "[0;32m"
    echo "****************************************************************"
    echo "**                                                            **"
    echo "**                       GIT PUSH REVIEW                      **"
    echo "**  1.input the name of branch                                **"
    echo "**  2.input the names of reviewers                            **"
    echo "****************************************************************"
	echo "[0m"
} 

D(){
    case $M in 
        branch_name)
            echo "Warning: You must enter a branch name"
            ;;
        reviewers)
            ;;
        *)
            ;;
    esac
}

P(){
    [ $# -eq 0 ] && exit 0
    echo -n "$*"
}

result(){
    local result="Push Review Failed!!"
    if [ $1 -eq 0 ]
	then	
	result="Push review Successfully!"
	echo "[0;32m"
	else
	echo "[0;31m"
    fi
	
    echo "+----------------------------Result----------------------------+"
    echo "|                                                              |"
    echo "    "$result
    echo "|                                                              |"
    echo "+--------------------------------------------------------------+"
	echo "[0m"
}


welcome

#input the branch name
M=branch_name
let W=W+1;
while [ 1 ]
do
    P $W "Please input the name of branch:" && read NAME
    [ x$NAME = x ] && D && continue
    break
done

#input the reviewers
let W=W+1;
M=reviewers

P $W "Please input the names of reviewers(eg:A,B,C):" && read REVIEWS
review=`echo $REVIEWS |sed 's/,/ /g'`
reviews=
count=0
for pa in $review
do
    [ x$pa = x ] && continue
    pa='r='$pa'@chinaotec.com'
    [ $count -gt 0 ] && reviews=$reviews,
    reviews=$reviews$pa
    count=1
done

#assme and do command
path=${GIT_REFERENCE_PATH}/${NAME}
push_cmd="git ${ACTION} ${PARM} ${path}"
[ x$reviews != x ] && push_cmd=$push_cmd'%'$reviews
echo $push_cmd

`$push_cmd` >/dev/null 2>&1
code=$?
result $code

exit 0

