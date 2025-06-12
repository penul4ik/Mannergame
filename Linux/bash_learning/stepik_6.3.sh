#!/bin/bash

usage() {
 cat << EOF
SINOPSIS: 
  $0 [OPTIONS] [DIR 1] [DIR 2] ...  [DIR N]

  A script for search valentinki in files in the directory.

OPTIONS:
  -h, --help                 Print this message.
  --max_age (int VALUE)      Filtering valentinki by max age.
  --min_age (int VALUE)      Filtering valentinki by min age.
  --min_income (int VALUE)   Filtering valentinki by min income.
  --min_height (int VALUE)   Filtering valentinki by min height.
  -i (int VALUE)             Girl number for compare
  [DIR 1]		     Specifies the directory for search files with valentinki and compare with DIR 2.
  [DIR 2]                    Compare valentinki with DIR 2 and print uniq valentinki. If not specified, it filters valentinki in DIR 1 and print them.
  [DIR N]                    

EXAMPLES: 
  $0 -h
  $0 --min_age 18 --min_income 100000 ~/valentinki ~/valentinki2
EOF
exit 1
}

OPTS=$(getopt -o h,i: --long help,max_age:,min_age:,min_income:,min_height:,i: -- "$@")

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
		-i | --i)
			GIRL=$2
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

if [ ! -d "$1" ]; then
	echo "$1 not a directory" >&2
	usage
elif [ ! -d "$2" ]; then
	echo "$2 not a directory" >&2
	usage
fi

cat $1/* > /tmp/tmp_filter.txt
grep -E 'Ты[[:space:]]мне[[:space:]]очень[[:space:]]нравишься[[:space:]]+[А-Яа-яЁё]{2,20}[[:space:]]+[0-9]{3}[[:space:]]+[1-9][0-9]{4,}[[:space:]]+[0-9]{2}' /tmp/tmp_filter.txt | grep -vi "иван" > /tmp/tmp_filter_1.txt
rm /tmp/tmp_filter.txt

cat $2/* > /tmp/tmp_filter.txt
grep -E 'Ты[[:space:]]мне[[:space:]]очень[[:space:]]нравишься[[:space:]]+[А-Яа-яЁё]{2,20}[[:space:]]+[0-9]{3}[[:space:]]+[1-9][0-9]{4,}[[:space:]]+[0-9]{2}' /tmp/tmp_filter.txt | grep -vi "иван" > /tmp/tmp_filter_2.txt
rm /tmp/tmp_filter.txt

grep -vFxf /tmp/tmp_filter_2.txt /tmp/tmp_filter_1.txt > /tmp/ffc.txt
rm /tmp/tmp_filter_1.txt /tmp/tmp_filter_2.txt

while read line; do 
	if [ $(echo $line | awk '{print$6}') -ge $MIN_HEIGHT ] && [ $(echo $line | awk '{print$7}') -ge $MIN_INCOME ] && [ $(echo $line | awk '{print$8}') -le $MAX_AGE ] && [ $(echo $line | awk '{print$8}') -ge $MIN_AGE ];then
	       echo $line
	fi
done < /tmp/ffc.txt
rm /tmp/ffc.txt

echo $GIRL
# Надо спарсить какую по номеру директорию указали в i и начать сравнивать с ней остальные указанные директории
# Все директории берем с помощью $@ и пихаем в массив из которого будем брать номер_девушки-1 директорию 
# И сравнивать с остальными элементами массива (директориями)
# Берем директорию для сравнения и берем кол-во оставшихся аргументов с директориями
# Запускаем цикл в котором сравниваем ФАЙЛ ИСХОДНЫЙ с файлами других девушек
# Если есть совпадения, удаляем их из исходного файла
# Выводим по итогу то, что осталось  
