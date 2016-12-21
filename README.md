# pikachu
brain fxckでピカチュウ語つくってみた
ピカチュウの歌 (<https://www.youtube.com/watch?v=9JUwoSUZA34>)が可愛すぎてピカチュウ語が喋りたくてたまらなくなったのでbrain fuckのピカチュ語バージョンをつくってみました。
#定義 
* `> → ピカ` 
  
  (ポインタをインクリメントする。ポインタをptrとすると、C言語の「ptr++;」に相当する。)  
* `< → ピ` 
  
  (ポインタをデクリメントする。C言語の「ptr--;」に相当。)  
* `+ → ピカチュー` 
  
  (ポインタが指す値をインクリメントする。C言語の「(*ptr)++;」に相当。)  
* `- → ピッ` 
  
  (ポインタが指す値をデクリメントする。C言語の「(*ptr)--;」に相当。)  
* `. → ピッカ` 
  
  (ポインタが指す値を出力に書き出す。C言語の「putchar(*ptr);」に相当。)  
* `, → ピーカー` 
  
  (入力から1バイト読み込んで、ポインタが指す先に代入する。C言語の「*ptr=getchar();」に相当。)  
* `[ → ピィ` 
  
  (ポインタが指す値が0なら、対応する ] の直後にジャンプする。C言語の「while(*ptr){」に相当。)  
* `] → ピカァ`
  
  (ポインタが指す値が0でないなら、対応する [ （の直後[1]）にジャンプする。C言語の「}」に相当。)  
  
#使い方
##encode
ピカチュウ語translatot ( <https://pikachu-lang.appspot.com/pikachu> )でできます。
GAEを使っています。コードは app/ 以下。

##decode
rubyのr-fxxkというライブラリを使うとすごく簡単に作れました。
```
ruby pikachu_lang.rb helloworld.pk
```

##brain fuck →　pikach
```
ruby bfToPika.rb helloworld.bf
```
