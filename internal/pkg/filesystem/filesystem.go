package filesystem

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

func PrintDiskInfo() error {
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Ошибка при получении информации о дисках:", err)
		return err
	}

	for _, partition := range partitions {
		fmt.Printf("Устройство: %s\n", partition.Device)
		fmt.Printf("Метка тома: %s\n", partition.Mountpoint)
		fmt.Printf("Тип файловой системы: %s\n", partition.Fstype)

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println("Ошибка при получении информации о использовании диска:", err)
		} else {
			fmt.Printf("Размер диска: %s\n", byteCountDecimal(usage.Total))
		}

		fmt.Println()
	}
	return nil
}

// Функция для преобразования байтов в человеко-читаемый формат
func byteCountDecimal(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
