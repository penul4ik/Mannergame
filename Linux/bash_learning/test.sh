#!/bin/bash -x

# Только длинные опции — никаких -o
PARSED=$(getopt --long help,verbose --name "$0" -- "$@")
if [[ 2 -ne 0 ]]; then
  echo "Ошибка при разборе аргументов"
  exit 1
fi

eval set -- ""

while true; do
  case "" in
    --help)
      echo "Показать справку"
      shift
      ;;
    --verbose)
      echo "Режим подробного вывода"
      shift
      ;;
    --)
      shift
      break
      ;;
    *)
      echo "Неизвестная опция: "
      exit 1
      ;;
  esac
done

echo "Оставшиеся аргументы: "
