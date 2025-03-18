# Football Player Performance Analytics

ระบบวิเคราะห์ประสิทธิภาพนักฟุตบอล เป็นระบบที่ช่วยในการจัดการและวิเคราะห์ข้อมูลประสิทธิภาพของนักฟุตบอล

## คุณสมบัติหลัก

- บันทึกและจัดการข้อมูลนักฟุตบอล
- เก็บสถิติการแข่งขันรายบุคคล
- วิเคราะห์ประสิทธิภาพผู้เล่น
- เปรียบเทียบสถิติระหว่างผู้เล่น
- สร้างรายงานพัฒนาการของผู้เล่น
- คาดการณ์ผลงานในอนาคต

## เทคโนโลยีที่ใช้

- Go 1.21+
- PostgreSQL
- Docker & Docker Compose
- Go Modules
- Testing (testify)

## การติดตั้งและการใช้งาน

### ความต้องการของระบบ

- Go 1.21 หรือสูงกว่า
- Docker และ Docker Compose
- PostgreSQL

### การติดตั้ง

1. Clone repository:
```bash
git clone https://github.com/yourusername/football-analyze.git
cd football-analyze
```

2. ติดตั้ง dependencies:
```bash
go mod download
```

3. สร้าง environment file:
```bash
cp .env.example .env
```

4. รัน application:
```bash
docker-compose up -d
go run cmd/api/main.go
```

## การทดสอบ

รันการทดสอบทั้งหมด:
```bash
go test ./...
```

## API Documentation

API documentation จะถูกสร้างโดยอัตโนมัติที่ `/docs` endpoint

## License

MIT License
