// +build windows

package file

import (
	"fmt"
	"golang.org/x/sys/windows"
	"sort"
	"syscall"
	"time"
)

// https://github.com/vinely/disks/blob/master/check_windows.go
// see https://docs.microsoft.com/en-us/windows/desktop/fileio/displaying-volume-paths
// or https://blog.csdn.net/hurricane_0x01/article/details/51516550

// VolumeInfo @todo
type VolumeInfo struct {
	MountPoint             string
	Name                   string
	SerialNumber           uint32
	MaximumComponentLength uint32
	FileSystemFlags        uint32
	FileSystemName         string
}

var volumeMap map[uint64]VolumeInfo

func init() {
	volumeMap = make(map[uint64]VolumeInfo)
}

const (
	// MaxVolumeLabelLength is the maximum number of characters in a volume label.
	MaxVolumeLabelLength = windows.MAX_PATH + 1

	// MaxVolumeNameLength is the maximum number of characters in a volume name.
	//
	//   \\?\Volume{xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}\
	MaxVolumeNameLength = windows.MAX_PATH + 1 // 50?

	// MaxFileSystemNameLength is the maximum number of characters in a file
	// system name.
	MaxFileSystemNameLength = windows.MAX_PATH + 1

	// MaximumComponentLength @todo
	MaximumComponentLength = 255 //for FAT.
)

var (
	// VolumeName @todo
	VolumeName [MaxVolumeNameLength]uint16
	// CheckEachTimeout check timeout for every volume
	CheckEachTimeout = time.Duration(5)
)

func loadVolumeMap() error {
	var (
		volumeMountPoint       [MaxFileSystemNameLength]uint16
		retLen                 uint32
		VolumeNameBuffer       = make([]uint16, syscall.MAX_PATH+1)
		nVolumeNameSize        = uint32(len(VolumeNameBuffer))
		VolumeSerialNumber     uint32
		MaximumComponentLength uint32
		FileSystemFlags        uint32
		FileSystemNameBuffer          = make([]uint16, 255)
		nFileSystemNameSize    uint32 = syscall.MAX_PATH + 1
	)

	t := time.NewTimer(CheckEachTimeout * time.Second)
	defer t.Stop()
	stopflag := make(chan bool, 1)
	goWrap := func() {
		sVolumeName := windows.UTF16ToString(VolumeName[:])
		err := windows.GetVolumePathNamesForVolumeName(&VolumeName[0], &volumeMountPoint[0], MaxFileSystemNameLength, &retLen)
		if err != nil {
			fmt.Printf("Failed to get volume information for %s: %s\n", sVolumeName, err.Error())
			return
		}
		mountPoint := windows.UTF16ToString(volumeMountPoint[:])
		if mountPoint == "" {
			return
		}

		err = windows.GetVolumeInformation(
			&volumeMountPoint[0],     // rootPathName *uint16
			&VolumeNameBuffer[0],     // volumeNameBuffer *uint16
			nVolumeNameSize,          // volumeNameSize uint32
			&VolumeSerialNumber,      // volumeNameSerialNumber *uint32
			&MaximumComponentLength,  //  maximumComponentLength *uint32
			&FileSystemFlags,         // fileSystemFlags *uint32
			&FileSystemNameBuffer[0], // fileSystemNameBuffer *uint16
			nFileSystemNameSize)      // fileSystemNameSize uint32
		if err != nil {
			fmt.Printf("Failed to get volume information for %s (%s): %s\n", mountPoint, sVolumeName, err.Error())
			return
		}
		vi := VolumeInfo{}
		vi.MountPoint = mountPoint
		vi.Name = syscall.UTF16ToString(VolumeNameBuffer)
		vi.SerialNumber = VolumeSerialNumber
		vi.MaximumComponentLength = MaximumComponentLength
		vi.FileSystemFlags = FileSystemFlags
		vi.FileSystemName = syscall.UTF16ToString(FileSystemNameBuffer)
		//fmt.Printf("vi=%+v\n", vi)
		volumeMap[uint64(VolumeSerialNumber)] = vi
		stopflag <- true
	}

	hvol, err := windows.FindFirstVolume(&VolumeName[0], MaxVolumeNameLength)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	defer windows.FindVolumeClose(hvol)

	go goWrap()
	select {
	case <-t.C:
	case <-stopflag:
	}

	for {
		if err := windows.FindNextVolume(hvol, &VolumeName[0], MaxVolumeNameLength); err != nil {
			break
		}
		t.Reset(CheckEachTimeout * time.Second)
		go goWrap()
		select {
		case <-t.C:
		case <-stopflag:
		}
	}

	return nil
}

// VolumeName @todo
func (f *File) VolumeName() (volume string, err error) {
	volume = fmt.Sprintf("%x", f.volumeID)
	if len(volumeMap) == 0 {
		err := loadVolumeMap()
		if err != nil {
			return volume, err
		}
	}
	vol, ok := volumeMap[f.volumeID]
	if ok {
		return vol.MountPoint, nil
	}
	return volume, nil
}

// GetVolumes @todo
func GetVolumes() ([]string, error) {
	var err error
	if len(volumeMap) == 0 {
		err = loadVolumeMap()
	}
	var volumes []string
	for _, volume := range volumeMap {
		volumes = append(volumes, volume.MountPoint)
	}
	sort.Strings(volumes)
	return volumes, err
}
