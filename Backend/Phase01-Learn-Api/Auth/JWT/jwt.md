## JWT

JSON Web Token, yang berarti token ini menggunakan JSON (Javascript Object Notation) berbentuk string panjang yang sangat random, lalu token ini memungkinkan kita untuk mengirimkan data yang dapat diverifikasi oleh dua pihak atau lebih.

## Bagaimana cara JWT bekerja ?
Seperti password jadi ketika users berhasil melakukan Login maka server akan memberikan sebuah Token. Nanti Token tersebut akan disimpan oleh users pada Local Storage atau Cookies Browser dan bila users ingin mengakses halaman halaman tertentu maka harus menyertakan token tersebut. Untuk itu users akan mengirim balik token yang dikasih diawal tadi sebagai bukti bila user ini, sudah melakukan login. Sekarang kita akan lihat struktur dasarnya Tokennya dimana terdiri dari tiga bagian yaitu yang pertama header lalu kedua bagian payloadnya atau datanya dan yang ketiga adalah bagian verify signature.

## Struktur JWT
### Header
Header biasanya terdiri dari dua bagian: jenis token, yaitu JWT, dan algoritma penandatanganan yang digunakan, seperti HMAC SHA256 atau RSA.

```bash
{
  "alg": "HS256",
  "typ": "JWT"
}
Payload
```

### Payload
sebagai infomasi atau data yang ingin kita kirimkan. Dalam penerapannya di otentikasi atau pun otorisasi, biasanya data ini berupa data yang sifatnya unik bagi user, seperti: email, id/uuid, dan juga data yang berkaitan dengan otorisasi seperti role, karena data tersebut akan digunakan sebagai tanda pengenal si pengirim token.

Oh ya, jangan pernah menyelipkan data yang sifatnya konfidental atau sensitif ke dalam JWT (seperti password).
```bash
{
  "sub": "1234567890",
  "name": "Minpo",
  "admin": true
}
```

### Verify Signature
adalah hasil dari Hash atau gabungan dari isi encode Header dan Payloadnya lalu ditambahkan kode secretnya. Signature ini berguna untuk memverifikasi bahwa header maupun payload yang ada dalam token tidak berubah dari nilai aslinya (karena untuk membuat payload dan header palsu itu cukup mudah).

Signature-nya sendiri tidak mungkin dapat diakali, karena sudah dalam berbentuk hash; yang mana adalah fungsi satu arah (tidak dapat dikembalikan ke nilai semula), dan meski kita tahu algoritma hashing-nya, kita juga memerlukan secret key yang mana hanya si pembuat aplikasi yang tahu.
```bash
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```
Jika kita satukan semuanya

Dan hasil ketiga bagian tersebut akan digabung dan otomatis di encode menjadi Token string random panjang seperti berikut

<img src="https://i.ibb.co/xHKBcf5/jwt.png"alt="JWT Result Example" />

Jika isi Header atau Payload dirubah maka isi Signature menjadi tidak valid. Satu hal lagi JWT ini tidak tergantung sama bahasa program tertentu jadi kita bisa mengimplementasikan di Laravel Codeigniter Node JS dan yang lainnya. Sekian dan Terimakasih.

Source <a href="https://jwt.io">JWT</a>
