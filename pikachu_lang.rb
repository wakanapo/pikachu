#!/usr/bin/env ruby
# coding: utf-8

require 'r-fxxk'

class Pikachu < Brainfuck
  nxt 'ピカ '
  inc 'ピカチュー '
  prv 'ピ '
  dec 'ピッ '
  opn 'ピィ '
  cls 'ピカァ '
  put 'ピッカ '
  get 'ピーカー '
end

puts Pikachu.new.fuck(ARGF.read)
