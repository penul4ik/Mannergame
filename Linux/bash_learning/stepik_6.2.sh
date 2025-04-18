#!/bin/bash

# Хотим фильтровать текст из файлов в указанной директории
# Принимаем на вход директорию, параметры фильтрации
# В указанной директории получаем список файлов с расширением *.txt, открываем каждый из них
# В соответствии с указанными параметрами применяем фильтр к тексту из открытого файла
# Отфильтрованный текст и путь к файлу с текстом выводим в stdout

usage() {
 cat << EOF
SINOPSIS: 
  $0 [OPTIONS] [DIR]

  A script for search valentinki in files in the directory.

OPTIONS:
  -h, --help                 Print this message.
  --max_age (int VALUE)      Filtering valentinki by max age.
  --min_age (int VALUE)      Filtering valentinki by min age.
  --min_income (int VALUE)   Filtering valentinki by min income.
  --min_height (int VALUE)   Filtering valentinki by min height.
  [DIR]                      Specifies the directory dor search files with valentinki. If empty that script will search valentinki in current directory.

EXAMPLES: 
  $0 -h
  $0 --min_age 18 --min_income 100000 ~/valentinki
EOF
exit 1
}

OPTS=$(getopt -o h --long help,max_age:,min_age:,min_income:,min_height: -n 'simple_getopt.sh' -- "$@")

if [ $? -ne 0 ]; then
	echo "failed to parse options with getops" >&2
	usage
fi

eval set -- "$OPTS"

HELP=false
MAX_AGE=200
MIN_AGE=0
MIN_INCOME=0
MIN_HEIGHT=0

while true; do
	case "$1" in
		-h | --help)
			usage
			shift
			;;
		--max_age)
			MAX_AGE=$2
			shift 2
			;;
		--min_age)
                        MIN_AGE=$2
                        shift 2
                        ;;
		--min_income)
                        MIN_INCOME=$2
                        shift 2
                        ;;
		--min_height)
			MIN_HEIGHT=$2
			shift 2
			;;
		--)
			shift
			break
			;;
		*)
			echo "Failed parse option" >&2
			usage
			;;
	esac
done

DIRECTORY="$@"

if [ ! -d $DIRECTORY ]; then
	echo "$DIRECTORY not a directory" >&2
	usage
fi

for file in $(ls -1 $DIRECTORY | grep -E '*.txt'); do
		valentinka=$(cat $file)
	if [[ "$valentinka" =~ Ты[[:space:]]мне[[:space:]]очень[[:space:]]нравишься[[:space:]]+[А-Яа-яЁё]{2,20}[[:space:]]+[0-9]{3}[[:space:]]+[1-9][0-9]{4,}[[:space:]]+[0-9]{2} && ! "$valentinka" =~ .*Иван.* && ! "$valentinka" =~ .*иван.* ]]; then
		AGE=$(echo $valentinka|awk '{print$8}')
		INCOME=$(echo $valentinka|awk '{print$7}')
		HEIGHT=$(echo $valentinka|awk '{print$6}')
		if [ "$AGE" -le "$MAX_AGE" ] && [ "$AGE" -ge "$MIN_AGE" ] && [ "$INCOME" -ge "$MIN_INCOME" ] && [ "$HEIGHT" -ge "$MIN_HEIGHT" ]; then
                	echo "$DIRECTORY$file$AGE"
			cat $file
		fi
	fi
done


echo "max_age $MAX_AGE, min_age $MIN_AGE, min_income $MIN_INCOME, min_height $MIN_HEIGHT, directory $DIRECTORY"


