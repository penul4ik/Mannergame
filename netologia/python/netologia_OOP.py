from datetime import datetime

class Account:
    def __init__(self, name, amount):
        self._name = name
        self.history = []
        if amount < 0:
            raise ValueError('Err: amount must be greater then null')
        else:
            self._amount = amount
            self.upd_history(f'Bank open for {self._name} with {self._amount} sum')

    def upd_history(self, msg: str):
        # timestamp: action
        msg = str(datetime.now().strftime('%Y-%m-%d')) + ': ' + msg
        self.history.append(msg)

    def plus_amount(self, num):
        if num < 0:
            raise ValueError('Err: num must be greater then null')
        else:
            self._amount += num
            self.upd_history(f'plus {num}')

    def minus_amount(self, num):
        if self._amount < num:
            raise ValueError('Err: amount less then num')
        else:
            self._amount -= num
            self.upd_history(f'minus {num}')

    def display_info(self, which_info):
        if which_info == 'history':
            return self.history
        elif which_info == 'amount':
            return self._amount
        else:
            print('history or amount?')

dima = Account('Dima', 100)
print(dima.display_info('amount'))
dima.plus_amount(100)
dima.minus_amount(50)
print(dima.display_info('amount'))
print(dima.display_info('history'))

max = Account('Max', 100)
print(max.display_info('amount'))
max.plus_amount(100)
max.minus_amount(50)
print(max.display_info('amount'))
print(max.display_info('history'))
