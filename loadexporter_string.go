// Code generated by "stringer -type=LoadExporter"; DO NOT EDIT

package prombench

import "fmt"

const _LoadExporter_name = "ExporterIncExporterStaticExporterRandCyclicExporterOscillate"

var _LoadExporter_index = [...]uint8{0, 11, 25, 43, 60}

func (i LoadExporter) String() string {
	if i < 0 || i >= LoadExporter(len(_LoadExporter_index)-1) {
		return fmt.Sprintf("LoadExporter(%d)", i)
	}
	return _LoadExporter_name[_LoadExporter_index[i]:_LoadExporter_index[i+1]]
}
