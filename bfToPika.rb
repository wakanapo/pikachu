#!/usr/bin/env ruby
# coding: utf-8

begin
  ARGF.each_char do |char|
    case char
    when '>' then
      print 'ピカ '
    when '<' then
      print 'ピ '
    when '+' then
      print 'ピカチュー '
    when '-' then
      print 'ピッ '
    when '.' then
      print 'ピッカ '
    when ',' then
      print 'ピーカー '
    when '[' then
      print 'ピィ '
    when ']' then
      print 'ピカァ '
    end
    
  end
  
end
