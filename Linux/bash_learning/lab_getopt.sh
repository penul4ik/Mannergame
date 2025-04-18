#!/bin/bash -x

OPTS=$(getopt -o h --long help,verbose -n 'simple_getopt.sh' -- "$@")

if [ $? -ne 0 ]; then
	echo "Failed to parse options" >&2
	exit 1
fi

eval set -- "$OPTS"

HELP=false
VERBOSE=false


while true; do
	case "$1" in
		-h | --help)
			HELP=true
			shift
			;;
		--verbose)
			VERBOSE=true
			shift
			;;
		--)
			shift
			break
			;;
		*)
			echo "Internal error!"
			exit 1
			;;
	esac
done


if [ "$HELP" = true ]; then
	echo "Help is enabled"
fi

if [ "$VERBOSE" = true ]; then
	echo "Verbose mode is enabled"
fi


echo "Remaining arguments: $@"
