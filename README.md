# gomatrix

## 概要

gomatrixは、Go言語で記述された行列演算ライブラリです。行列の作成、加算、減算、スカラー倍、シグモイド関数適用、転置などの基本的な行列演算を提供します。

## 特徴

* シンプルで使いやすいAPI
* 行列演算の基本的な機能を提供
* Go言語で記述されており、軽量で高速

## インストール

```bash
go get github.com/takoyaki-3/gomatrix 
```

## 使い方

### 行列の作成

```Go
// 10x5の行列を作成
a := mat.NewMatrix(10, 5)

// -0.5から0.5のランダムな値で初期化された5x10の行列を作成
b := mat.NewRandom05to05Matrix(5, 10)
```

### 行列演算

```Go
// 行列にスカラー値を加算
a.ScalarPlus(1)

// 行列にシグモイド関数を適用
a.Sigmoid()

// 行列の積を計算 (a x b)
c := mat.NewDotMatrix(a, b)

// 行列の転置を取得
t := m.GetT()
```

### その他の関数

* `Print()`: 行列の内容を標準出力に出力
* `Put(i, j int, x float32)`: 行列の(i, j)要素に値xを設定
* `Get(i, j int) float32`: 行列の(i, j)要素の値を取得
* `Multiplication(a *Matrix) bool`: 行列の要素同士の積を計算
* `Scalar(k float32)`: 行列にスカラー値を乗算
* `ScalarDivision(k float32)`: 行列をスカラー値で除算
* `Plus(mm *Matrix)`: 行列を加算
* `Minus(mm *Matrix)`: 行列を減算
* `NewCopiedMatrix(mm *Matrix) *Matrix`: 行列のコピーを作成
* `NewCopyMatrix(a *Matrix) *Matrix`: 行列のコピーを作成 (Plus()を使用)


## ファイルツリー

```
gomatrix/
├── go.mod
├── model.go
└── sample/
    └── main.go
```

### ファイルの説明

* **go.mod:** Goモジュールの依存関係を管理するファイルです。
* **model.go:** 行列構造体 `Matrix` と関連する関数を定義しています。
* **sample/main.go:** `gomatrix` ライブラリの使用方法を示すサンプルコードです。


## 実行例

`sample/main.go` を実行するには、以下のコマンドを実行します。

```bash
cd sample
go run main.go
```
