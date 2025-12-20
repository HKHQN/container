package main

import (
    "fmt"
    "time"
)

// Định nghĩa các trạng thái logistics
type Status string

const (
    OrderPlaced           Status = "Đặt hàng"
    Packed                Status = "Đóng gói"
    InTransitFromWarehouse Status = "Đang vận chuyển từ kho"
    ArrivedSortingCenter  Status = "Đến trung tâm phân phối"
    OutForDelivery        Status = "Đang giao hàng"
    Delivered             Status = "Đã giao thành công"
    FailedDelivery        Status = "Giao hàng thất bại"
)

// Cấu trúc lưu thông tin theo dõi lô hàng
type TrackingUpdate struct {
    Status    Status
    Timestamp time.Time
    Location  string
    Note      string
}

// Cấu trúc lô hàng
type Shipment struct {
    TrackingNumber string
    Updates        []TrackingUpdate
}

// Thêm cập nhật trạng thái với thời gian Việt Nam
func (s *Shipment) AddUpdate(status Status, location, note string) {
    // Load múi giờ Việt Nam
    vietnamLoc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
    if err != nil {
        panic(err) // Trong thực tế nên handle error tốt hơn
    }

    update := TrackingUpdate{
        Status:    status,
        Timestamp: time.Now().In(vietnamLoc),
        Location:  location,
        Note:      note,
    }
    s.Updates = append(s.Updates, update)
}

// In lịch sử theo dõi
func (s *Shipment) PrintTrackingHistory() {
    fmt.Printf("Mã vận đơn: %s\n", s.TrackingNumber)
    fmt.Println("Lịch sử theo dõi:")
    fmt.Println("--------------------------------------------------")
    for i, update := range s.Updates {
        fmt.Printf("%d. [%s]\n", i+1, update.Timestamp.Format("02/01/2006 15:04:05"))
        fmt.Printf("   Trạng thái: %s\n", update.Status)
        fmt.Printf("   Địa điểm: %s\n", update.Location)
        if update.Note != "" {
            fmt.Printf("   Ghi chú: %s\n", update.Note)
        }
        fmt.Println("--------------------------------------------------")
    }
}

func main() {
    // Tạo lô hàng mới
    shipment := Shipment{
        TrackingNumber: "VN123456789XYZ",
    }

    // Mô phỏng quá trình logistics với các mốc thời gian cách nhau
    shipment.AddUpdate(OrderPlaced, "Hà Nội", "Khách hàng đặt hàng online")
    time.Sleep(2 * time.Second) // Giả lập thời gian trôi qua

    shipment.AddUpdate(Packed, "Kho Hà Nội", "Đơn hàng đã được đóng gói")
    time.Sleep(1 * time.Second)

    shipment.AddUpdate(InTransitFromWarehouse, "Kho Hà Nội → TP.HCM", "Đang trên đường vận chuyển")
    time.Sleep(3 * time.Second)

    shipment.AddUpdate(ArrivedSortingCenter, "Trung tâm phân phối TP.HCM", "")
    time.Sleep(1 * time.Second)

    shipment.AddUpdate(OutForDelivery, "Quận 1, TP.HCM", "Shipper đang giao hàng")
    time.Sleep(2 * time.Second)

    shipment.AddUpdate(Delivered, "Quận 1, TP.HCM", "Khách hàng đã nhận hàng thành công")

    // In ra lịch sử theo dõi
    shipment.PrintTrackingHistory()
}
