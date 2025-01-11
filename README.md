# pjsekai-overlay-ipad / プロセカ風動画作成補助ツール

pjsekai-overlay-ipad は、pjsekai-overlayをipad用に改造したものです

## 必須事項

- [AviUtl](http://spring-fragrance.mints.ne.jp/aviutl/) + [拡張編集プラグイン](http://spring-fragrance.mints.ne.jp/aviutl/) （[導入方法](https://aviutl.info/dl-innsuto-ru/)）
  （強く推奨：[patch.aul](https://scrapbox.io/ePi5131/patch.aul)）
- [AVIUtl_Unmult](https://github.com/mes51/AVIUtl_Unmult)
- AviUtlの基本的な知識

## 動画の作り方

1. [譜面を作る](https://wiki.purplepalette.net/create-charts)
2. [Sonolus](https://sonolus.com/)で譜面を撮影する
   - [Potato Leaves](https://github.com/sevenc-nanashi/potato_leaves)、または [Chart Cyanvas](https://cc.sevenc7c.com)で撮影してください。
   - 「Hide UI」をオンにしてください。
3. 撮影したプレイ動画のファイルをパソコンに転送する
   - Google Drive など
4. [ffmpeg](https://www.ffmpeg.org/)で再エンコードする
   - AviUtl で読み込むため
5. 下の利用方法に従って UI を後付けする

## 利用方法

0. 1280x720, 60fps で aviutl のプロジェクトを作成する
1. 注意　必ず、オブジェクトファイルから作成ではなくプロジェクトを作成してから、
2.  オブジェクトファイルをインポートしてください
3. 右の Releases から最新のバージョンの zip をダウンロードする
4. zip を解凍する
5. AviUtl を起動する
   - pjsekai-overlay-ipad が起動する前に AviUtl を起動するとオブジェクトのインストールが行われます。
6. `pjsekai-overlay-ipad.exe` を起動する
7. 譜面 ID を入力する
   - Potato Leaves の場合は `ptlv-` を、Chart Cyanvas の場合は `chcy-` を先頭につけたまま入力してください。
8. 動画を出力したら、形式をmp4などの形式からmovにします。
9. iphoneやipadに完成した動画を送ったら、写真アプリに保存して、動画を編集する画面に行ってください、
10.切り取りに移動して、右上から3番目のボタンを押します。そしたら、下の4:3を選択して保存します。

## 利用規約

動画の概要欄などに、

プロセカ風動画作成補助ツール：
  オリジナル版 https://github.com/sevenc-nanashi/pjsekai-overlay
  作成：名無し｡ （ https://sevenc7c.com ）
  ipad版：https://github.com/Nsan736/pjsekai-overlay-ipad
  作成：Nsan  (https://youtube.com/channel/UCOyHw_H6jFwprHaPuvzDLDQ?si=pHFqqH5NuMxhuAh8)
  を入れてください

