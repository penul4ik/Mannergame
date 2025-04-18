#!/bin/bash

echo "Оригинальные аргументы:"
for arg in "$@"; do
  echo "[$arg]"
done
echo

# Разбор аргументов
PARSED=$(getopt -o a: -- "$@")
echo "Результат от getopt:"
echo "$PARSED"
echo

#########################
# БЕЗ eval
#########################
set -- "$PARSED"

echo "❌ Без eval set --:"
i=1
for arg in "$@"; do
  echo "[$i] $arg"
  ((i++))
done
echo

#########################
# С eval
#########################
PARSED=$(getopt -o a: -- "$@")  # ещё раз для чистоты эксперимента
eval set -- "$PARSED"

echo "✅ С eval set --:"
i=1
for arg in "$@"; do
  echo "[$i] $arg"
  ((i++))
done

