# Golang Clean Architecture - HATEOAS, Pagination, Unit Tested, Cassandra Rest API

Bu proje, modern yazılım geliştirme prensiplerini temel alarak Golang kullanılarak geliştirilmiş bir REST API çözümüdür. **Clean Architecture** ilkelerine uygun olarak tasarlanmıştır ve hiper medya kontrolleri (HATEOAS), sayfalama (pagination), birim testleri ile doğrulama ve Cassandra veritabanı entegrasyonunu desteklemektedir. Dağıtılabilirliği, performansı ve sürdürülebilirliği ön planda tutan bu proje, yazılım geliştirme süreçlerinde bir referans noktası olmayı hedeflemektedir.

## Özellikler

- Temiz Mimari (Clean Architecture) sayesinde bağımlılıklar minimuma indirilmiştir ve kodun test edilebilirliği artırılmıştır.
- HATEOAS desteği sayesinde REST API aracılığıyla hiper medya kontrolleri sunularak kullanıcı deneyimi geliştirilmiştir.
- Büyük veri kümelerinde hızlı ve verimli veri görüntüleme için optimize edilmiş sayfalama mekanizması bulunmaktadır.
- Birim testleri ile kod doğruluğu ve güvenilirliği sağlanmıştır.
- Cassandra veritabanı ile entegre edilen yüksek performanslı veri yönetimi imkanı sunulmaktadır.

## Kullanılan Teknolojiler

Golang, Cassandra, Docker (opsiyonel), HATEOAS ve Postman gibi araçlar projede yer almaktadır. Bu araçlar sayesinde hem geliştirme süreci optimize edilmiş hem de test ve entegrasyon süreçleri kolaylaştırılmıştır.

## Proje Yapısı

Projede kullanılan dosya ve dizinlerin genel yapısı şu şekildedir:
- **cmd/**: Uygulamanın giriş noktası.
- **internal/**: İş mantığı, kullanım senaryoları, veri erişim katmanı ve API gibi modüllerin bulunduğu ana dizin.
- **configs/**: Proje yapılandırma dosyaları.
- **tests/unit/**: Birim testlerini içeren dizin.

## Kurulum ve Çalıştırma

Projeyi klonladıktan sonra, gerekli bağımlılıkları yüklemek için `go mod tidy` komutunu çalıştırın. Cassandra bağlantısı için `config/config.yaml` dosyasını düzenleyin ve ardından `go run cmd/main.go` komutu ile uygulamayı başlatabilirsiniz.

## API Uç Noktaları

Proje aşağıdaki uç noktaları desteklemektedir:
- **GET /resources**: Sayfalama destekli veri listesi.
- **POST /resources**: Yeni bir kaynak oluşturma.
- **PATCH /resources/{id}**: Mevcut bir kaynağı güncelleme.
- **DELETE /resources/{id}**: Bir kaynağı silme.

## Testler

Birim testlerini çalıştırmak için `go test ./...` komutunu kullanabilirsiniz. Kapsamlı test senaryoları sayesinde kod kalitesi sağlanmaktadır.

## Katkıda Bulunma

Projeye katkıda bulunmak için:
1. Projeyi forklayın.
2. Yeni bir özellik dalı oluşturun.
3. Değişikliklerinizi commit'leyin.
4. Pull request göndererek katkılarınızı paylaşın.

## Lisans

Bu proje [MIT Lisansı](LICENSE) altında lisanslanmıştır.

## İletişim

Proje ile ilgili sorularınız ya da önerileriniz için iletişim kurabilirsiniz:
- E-posta: example@domain.com
- GitHub: [kullanıcı_adı](https://github.com/kullanıcı_adı)

