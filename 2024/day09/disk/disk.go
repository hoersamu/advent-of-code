package disk

import (
	"common"
	"slices"
)

type Disk struct {
	Blocks []DiskBlock
}

func NewDisk(diskMap string) *Disk {
	blocks := []DiskBlock{}
	mapInt := common.SplitAndTransform(diskMap, "", common.MustAtoi)

	for i, length := range mapInt {
		fileId := -1
		if common.IsEven(i) {
			fileId = i / 2
		}

		blocks = append(blocks, DiskBlock{FileId: fileId, Length: length})
	}

	return &Disk{Blocks: blocks}
}

func (d *Disk) getLastFileIndex() int {
	for i := len(d.Blocks) - 1; i >= 0; i-- {
		if !d.Blocks[i].IsFree() {
			return i
		}
	}
	return -1
}

func (d *Disk) moveBlock(destIndex, sourceIndex int) {
	sourceBlock := d.Blocks[sourceIndex]
	destBlock := d.Blocks[destIndex]

	if !destBlock.IsFree() {
		panic("dest block is not free")
	}

	sourceLength, destLength := sourceBlock.Length, destBlock.Length
	if sourceLength <= destLength {
		d.Blocks[destIndex] = sourceBlock.Copy()

		d.Blocks[sourceIndex].FileId = -1
		if sourceLength < destLength {
			d.Blocks = slices.Insert(d.Blocks, destIndex+1, DiskBlock{FileId: -1, Length: destLength - sourceLength})
		}
	} else {
		d.Blocks[destIndex].FileId = sourceBlock.FileId
		d.Blocks[sourceIndex].Length = sourceLength - destLength
	}
}

func (d *Disk) DefragmentHard() {
	for i := 0; i < len(d.Blocks); i++ {
		if !d.Blocks[i].IsFree() {
			continue
		}

		lastFileIndex := d.getLastFileIndex()
		if i > lastFileIndex {
			break
		}

		d.moveBlock(i, lastFileIndex)
	}
}

func (d *Disk) DefragmentSoft() {
	for i := len(d.Blocks) - 1; i >= 2; i-- {
		if d.Blocks[i].IsFree() {
			continue
		}

		blockLength := d.Blocks[i].Length

		for j := 0; j < i; j++ {
			blockToCheck := &d.Blocks[j]

			if !blockToCheck.IsFree() || blockLength > blockToCheck.Length {
				continue
			}

			length := len(d.Blocks)

			d.moveBlock(j, i)

			if len(d.Blocks) >= length {
				i++
			}
			break
		}
	}
}

func (d *Disk) CalculateChecksum() int {
	checksum := 0
	index := 0
	for _, block := range d.Blocks {
		if block.IsFree() {
			index += block.Length
			continue
		}
		for range block.Length {
			checksum += block.FileId * index
			index++
		}
	}
	return checksum
}

type DiskBlock struct {
	FileId int
	Length int
}

func (d *DiskBlock) IsFree() bool {
	return d.FileId == -1
}

func (d *DiskBlock) Copy() DiskBlock {
	return DiskBlock{
		FileId: d.FileId,
		Length: d.Length,
	}
}
