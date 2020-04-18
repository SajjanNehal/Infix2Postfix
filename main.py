#!/usr/bin/python3
import sys
from stack import Stack

postfix = ''
opstack = Stack()

def PopStackUntilEmpty():
	global postfix
	while not opstack.empty():
		postfix += opstack.peek()
		opstack.pop()

def PopStackUntilHigherPres(nextop):
	global postfix
	while not IsPresHigh(nextop, opstack.peek()):
		postfix += opstack.peek()
		opstack.pop()

def GetOperatorWeight(op):
	weight = -1
	if op == "+" or op == "-":
		weight = 1
	elif op == "*" or op == "/":
		weight = 2
	elif op == "^":
		weight = 3
	return weight


def IsPresHigh(nextchar, stacktop):
	nextchar_weight = GetOperatorWeight(nextchar)
	stacktop_weight = GetOperatorWeight(stacktop)

	if nextchar_weight == stacktop_weight:
		return True
	elif nextchar_weight > stacktop_weight:
		return True
	else:
		return False

def PopStackUntilLeftParenthesis():
	global postfix
	while opstack.peek() != "(":
		postfix += opstack.peek()
		opstack.pop()
	opstack.pop()

def IsOperator(C):
	operators = ["+", "-", "*", "/"]
	if any(oper in C for oper in operators):
		return True
	else:
		return False

def IterateString(input_infix):
	global postfix
	for i in input_infix:
		if IsOperator(i):
			if opstack.empty() or opstack.peek() == "(" or IsPresHigh(i, opstack.peek()):
				opstack.push(i)
			elif not IsPresHigh(i, opstack.peek()):
				PopStackUntilHigherPres(i)
		elif i == "(":
			opstack.push(i)
		elif i == ")":
			opstack.push(i)
		else:
			postfix += i
	PopStackUntilEmpty()
	return postfix

user_input = sys.argv[1]
print(IterateString(user_input))