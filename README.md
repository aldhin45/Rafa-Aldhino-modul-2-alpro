SOAL 2A NOMOR 1
Deklarasi Variabel:
satu, dua, dan tiga: Tiga variabel string digunakan untuk menyimpan input dari pengguna.
temp: Variabel string sementara digunakan untuk menukar nilai antara variabel-variabel lainnya.
 Nilai dari variabel satu disimpan sementara ke dalam variabel temp. Kemudian, nilai dua dipindahkan ke satu, nilai tiga dipindahkan ke dua, dan akhirnya nilai temp (yang semula adalah nilai satu) dipindahkan ke tiga.Program mencetak kembali ketiga string setelah dilakukan penukaran nilai. Urutan string akan berbeda dari output awal.Kode ini meminta pengguna untuk memasukkan tiga buah string, kemudian menukar urutan ketiga string tersebut, dan akhirnya mencetak hasil akhir setelah penukaran.

SOAL 2A NOMOR 2
Fungsi isLeapYear:
Menerima input berupa tahun (integer).
Mengembalikan true jika tahun tersebut adalah kabisat, false jika bukan.
Menggunakan aturan tahun kabisat:Tahun yang habis dibagi 400 adalah kabisat.Tahun yang habis dibagi 4 tetapi tidak habis dibagi 100 adalah kabisat. 
Fungsi main:
Menginisialisasi variabel year untuk menyimpan input tahun.
Meminta pengguna untuk memasukkan tahun.Memanggil fungsi tahunKabisat untuk memeriksa apakah tahun tersebut kabisat.
Menampilkan hasil pemeriksaan ke layar.

SOAL 2A NOMOR 3
Deklarasi Variabel:
radius: Menyimpan nilai jari-jari bola yang diinputkan oleh pengguna.
volume: Menyimpan hasil perhitungan volume bola.
surfaceArea: Menyimpan hasil perhitungan luas permukaan bola.
Perhitungan:
Volume: Menggunakan rumus volume bola (4/3 * π * r³). Nilai math.Pi memberikan nilai pi yang akurat, dan math.Pow digunakan untuk menghitung pangkat tiga dari jari-jari.
Luas Permukaan: Menggunakan rumus luas permukaan bola (4 * π * r²).

SOAL 2A NOMOR 4
Deklarasi Variabel:
celsius: Menyimpan nilai suhu dalam derajat Celsius yang diinputkan oleh pengguna.
fahrenheit, reamur, dan kelvin: Menyimpan hasil konversi suhu ke masing-masing satuan.
Perhitungan Konversi:
Menggunakan rumus-rumus konversi yang sudah baku untuk mengubah suhu dari Celsius ke Fahrenheit, Reamur, dan Kelvin

SOAL 2A NOMOR 5
Membaca Data:
Menggunakan loop for untuk meminta pengguna memasukkan 5 integer dan 3 karakter.
Menggunakan fmt.Scan dan fmt.Scanf untuk membaca input dari pengguna.
Menampilkan Hasil:
Menggunakan loop for untuk mencetak karakter yang telah diubah.
Menggunakan fmt.Printf untuk mencetak setiap karakter dengan format yang sesuai.

SOAL 2B NOMOR 1
Deklarasi Variabel:
warna1, warna2, warna3, warna4: Variabel string untuk menyimpan input warna dari pengguna pada setiap percobaan.
berhasil: Variabel boolean untuk menandai apakah semua percobaan berhasil atau tidak.
Loop Percobaan:
Menggunakan loop for untuk melakukan 5 kali percobaan.
Pada setiap percobaan, pengguna diminta untuk memasukkan 4 warna.
emeriksaan Kondisi:
Menggunakan if untuk memeriksa apakah urutan warna yang diinputkan sesuai dengan pola yang ditentukan.
Jika ada satu saja warna yang tidak sesuai, maka nilai berhasil diubah menjadi false dan loop dihentikan menggunakan break.

SOAL 2B NOMOR 2
Deklarasi Variabel:
N: Menyimpan jumlah maksimal bunga yang ingin dimasukkan.
bunga: Variabel sementara untuk menyimpan nama bunga pada setiap iterasi.
pita: Variabel string untuk menyimpan semua nama bunga yang dipisahkan oleh tanda "-"
Loop Input Bunga:
Menggunakan loop for untuk mengulang proses input bunga sebanyak N kali.
Pada setiap iterasi, pengguna diminta memasukkan nama bunga.
Jika pengguna memasukkan "SELESAI", loop akan dihentikan.
Nama bunga yang dimasukkan akan ditambahkan ke variabel pita.
Output:
Program mencetak semua nama bunga yang telah dimasukkan dalam variabel pita.
Program juga menghitung jumlah bunga dengan membagi pita berdasarkan tanda "-" dan menghitung panjangnya.

SOAL 2B NOMOR 3
Deklarasi Variabel:
beratKiri dan beratKanan: Menyimpan nilai berat belanjaan di masing-masing kantong.
selisih: Menyimpan nilai selisih berat antara kedua kantong.
Input Berat Belanjaan:
Pengguna diminta untuk memasukkan berat belanjaan di kedua kantong.
Hitung Selisih Berat:
Menghitung selisih berat antara kedua kantong menggunakan fungsi math.Abs untuk mendapatkan nilai absolut.

SOAL 2B NOMOR 4
Deklarasi Variabel:
k: Menyimpan nilai k yang diinputkan oleh pengguna.
fk: Menyimpan hasil perhitungan nilai f(k).
Perhitungan f(k):
Menggunakan rumus f(k) yang telah disebutkan di atas, nilai fk dihitung dengan memanfaatkan fungsi math.Pow untuk menghitung pangkat dua dari (4k+2).

SOAL 2C NOMOR 1
Deklarasi Variabel:
gram: Menyimpan berat parsel dalam gram yang diinputkan pengguna.
beratkg: Menyimpan bagian kilogram dari berat parsel.
beratGram: Menyimpan bagian gram dari berat parsel.
biayakg: Menyimpan biaya untuk bagian kilogram.
biayagram: Menyimpan biaya untuk bagian gram.
Konversi Berat:
Berat parsel dibagi 1000 untuk mendapatkan bagian kilogram (beratkg).
Sisa pembagian digunakan untuk mendapatkan bagian gram (beratGram).
Perhitungan Biaya:
Biaya untuk bagian kilogram dihitung dengan mengalikan beratkg dengan 10000.
Biaya untuk bagian gram dihitung berdasarkan kondisi:
Jika beratGram lebih dari 500, maka biaya per gram adalah 5.
Jika tidak, biaya per gram adalah 15.

SOAL 2C NOMOR 2
Deklarasi Variabel:
nam: Variabel bertipe float64 untuk menyimpan nilai akhir yang diinputkan pengguna.
nilaiAkhir: Variabel bertipe string untuk menyimpan nilai akhir dalam bentuk huruf.
Percabangan (if-else):
Program melakukan pengecekan terhadap nilai nam menggunakan serangkaian kondisi if-else.
Setiap kondisi memeriksa apakah nilai nam berada dalam rentang nilai tertentu.
Jika kondisi terpenuhi, maka nilai nilaiAkhir akan diisi dengan huruf yang sesuai.

SOAL 2C NOMOR 3
Deklarasi Variabel:
bilangan: Menyimpan bilangan yang akan dicari faktornya.
i: Variabel pengontrol dalam perulangan for untuk memeriksa setiap bilangan dari 1 hingga bilangan itu sendiri.
Mencari Faktor:
Perulangan for: Melakukan perulangan dari 1 hingga bilangan yang dimasukkan.
Kondisi if: Memeriksa apakah bilangan yang dimasukkan habis dibagi oleh i. Jika habis dibagi, maka i adalah faktor dari bilangan tersebut dan akan dicetak.
