# Interview K'Donn

### System Requirements
```json
{
    "golang": "^1.17",
}
```

### System Setup
```bash
# Run command
$ go mod vendor
```

### Run
```bash
$ go run main.go

# Open your browser and go to http://localhost:8080/
# You should see the product list
```

### Your Quest
เราต้องการเปลี่ยนการอ่านข้อมูลรายการสินค้าใหม่
(ปัจจุบันเราใช้การอ่านข้อมูลจาก array ที่มีข้อมูลสินค้าทั้งหมด)


เราต้องการให้คุณสร้างฟังก์ชันเพิ่อเปลี่ยนการอ่านข้อมูลรายการสินค้าใหม่
โดยเลือกระหว่าง 2 วิธี หรือทั้งสองวิธีก็สามารถทำได้
- MySQL
- MongoDB (จะพิจารณาเป็นพิเศษ)

บางครั้งระบบเก็บ Cache ของเราอาจจะไม่เพียงพอกับการทำงานของระบบ
เราจึงอยากให้เปลี่ยนจาก `json file` เป็น `redis` หรือ `memcached`
(Optional)


### หมายเหตุ
ห้ามแก้ไขโค้ดเดิม ยกเว้น
* [/service/product.go](service/product.go)

แต่คุณสามารถเพิ่มฟังก์ชันใหม่ หรือ class ใหม่ได้ตามความต้องการ
ซึ่งใน main.go ต้องเรียกใช้งานผ่าน [ProductService](service/product.go) เท่านั้น


### ส่งงาน
หลังจากท่านทำเสร็จแล้ว สามารถส่งงานของท่านได้ที่ Pull Request ของ Repository นี้ได้เลย

