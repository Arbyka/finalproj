# Gunakan image Golang resmi
FROM golang:1.21-alpine

# Set direktori kerja dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Ambil dependency
RUN go mod tidy

# Build binary aplikasi
RUN go build -o main .

# Jalankan aplikasi saat container di-start
CMD ["./main"]