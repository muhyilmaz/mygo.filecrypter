# mygo.filecrypter
# 1. Analiz:
## 1.1 Proje nedir?

mygo.filecrypter dosya şifreleme ve şifre çözümleme işlemini yapan bir projedir. Proje hedefinde, istenilen bir dosya girdi olarak alınacak ve şifreleme veya çözümlemesi yapılacaktır. 

## 1.2 Proje kapsamı

İlk etapta PDF dosyaların işlenmesi söz konusu olacaktır. PDF dosyalardan sonra Microsoft Office dosyaları işlenmesi hedeflenecektir. Şifreleme için kullanılacak metotlar AES-256 ve eliptic curve olarak düşünülmüştür. 

## 2. Yol haritası

## 2.1 Metodoloji:

Proje 11.12.2023 tarihinde başlamıştır. Her bir ay 4 haftalık parçalara bölümlenmiştir. Bu parçalar sprint olarak anılmaktadır. Her ay için hedefler çıkarılmıştır ve ilerleyen süreçte çıkarılacaktır. Ayrıca hafta bitiminde işlerin durumu raporlanacaktır. 

Yapılan her iş, yapım süresi ve bitiş tarihi olarak belirtilecektir. Hedeflenen gün ise iş başlangıcında verilecektir. 

## 2.2 Zaman planı:
Aşağıdaki alt başlıklarda proje başlangıç tarihinden itibaren sprintler yazılmıştır. 

### 2.2.1 Aralık 1. Hafta sprint hedefi:

#### Aralık 2. Hafta: ~~(11 - 17 Aralık)~~ (18 - 24 Aralık)
1. [ ] Şifreleme metotlarının tespiti.
   1. Hedef gün: 1 gün.
   2. Yapım süresi: ?
   3. Bulgular: 
      1. AES (Advanced Encryption Standard), verileri şifrelemek ve şifresini çözmek için kullanılan bir simetrik şifreleme algoritmasıdır. Bu algoritma, özellikle hassas bilgilerin güvenliğini sağlamak amacıyla tasarlanmıştır. AES, 128, 192 ve 256 bit olmak üzere üç farklı anahtar uzunluğu ile kullanılabilir.
      2. Simetrik şifreleme: AES, aynı anahtarın hem şifreleme hem de şifre çözme işlemlerinde kullanıldığı bir simetrik şifreleme algoritmasıdır. Bu, iletişim taraflarının aynı anahtarı paylaşması gerektiği anlamına gelir.
      3. AES, 128, 192 ve 256 bit anahtar uzunluklarıyla kullanılabilir. Anahtar uzunluğu arttıkça genel güvenlik artar, ancak işlemler daha yavaş hale gelir.
      4. AES, veriyi belirli boyutlardaki bloklar halinde şifreler. Standart olarak 128 bitlik bloklar kullanılır. Şifreleme işlemi her blok için ayrı ayrı gerçekleştirilir.

2. [X] Metin şifreleme yapılması
   1. Hedef gün: 2 gün.
   2. Yapım süresi: ?
   3. Bulgular: 
      1. Metin şifreleme için crpyto kütüphanesi kullanıldı.
      2. @huseyinozsoy katkıları ile bu taraf 21.12.2023 de test olarak tamamlanmıştır. 
3. [ ] Metin şifre çözümleme yapılması
   1. Hedef gün: 1 gün.
   2. Yapım süresi: ?
4. [ ] AES-256 ve eliptic curve algoritmalarının doğru çalıştırılması
   1. Hedef gün: 3 gün.
   2. Yapım süresi: ?


#### Aralık 3. hafta: ~~(18 - 24 Aralık)~~ (25 - 31 Aralık) (!Gün sayısı planlanacak)
1. [ ] Go ile dosya manipülasyonları, okuma, yazma
2. [ ] Dosya şifreleme denemeleri


