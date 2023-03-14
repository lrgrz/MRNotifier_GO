//go:build linux

package tray

var mergeIcon []byte = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x80,
	0x08, 0x06, 0x00, 0x00, 0x00, 0xc3, 0x3e, 0x61, 0xcb, 0x00, 0x00, 0x00,
	0x04, 0x73, 0x42, 0x49, 0x54, 0x08, 0x08, 0x08, 0x08, 0x7c, 0x08, 0x64,
	0x88, 0x00, 0x00, 0x07, 0x0b, 0x49, 0x44, 0x41, 0x54, 0x78, 0x9c, 0xed,
	0xdd, 0x5f, 0x88, 0x5c, 0x57, 0x1d, 0x07, 0xf0, 0xef, 0xef, 0x8c, 0x93,
	0x9b, 0x49, 0xb3, 0x6d, 0x4d, 0x45, 0x8d, 0x60, 0xff, 0x50, 0x75, 0xb1,
	0x5a, 0xa8, 0xb6, 0xc5, 0x56, 0x05, 0x6d, 0x28, 0x68, 0x4a, 0xdb, 0x97,
	0x42, 0x04, 0x45, 0xec, 0x5a, 0xcc, 0x4a, 0x49, 0x66, 0x5c, 0x06, 0x52,
	0x7c, 0xf0, 0x21, 0x08, 0x42, 0xb2, 0x8e, 0x93, 0x7b, 0xef, 0xc4, 0xb0,
	0xa2, 0xee, 0x83, 0xa9, 0x0f, 0x8b, 0x60, 0xf1, 0xa1, 0x42, 0x2b, 0x16,
	0xfa, 0x54, 0xd0, 0x50, 0x68, 0xad, 0xa1, 0x4f, 0xa5, 0xfe, 0x81, 0x04,
	0x12, 0xd0, 0x5d, 0x25, 0x33, 0x9d, 0xdc, 0xf3, 0xf3, 0xc1, 0x14, 0xda,
	0x2b, 0x26, 0x7b, 0xef, 0xfc, 0xce, 0x99, 0x39, 0x73, 0x7e, 0x9f, 0xc7,
	0xdd, 0x39, 0xbf, 0x73, 0xd8, 0xfb, 0xdd, 0x7b, 0xe6, 0xde, 0xdf, 0xdd,
	0x1d, 0x40, 0x29, 0xa5, 0x94, 0x52, 0x4a, 0x29, 0xa5, 0x94, 0x52, 0x4a,
	0x29, 0xa5, 0x94, 0x52, 0xf3, 0x8c, 0xa6, 0xbd, 0x80, 0x59, 0x93, 0xa6,
	0xe9, 0xdd, 0x00, 0x96, 0x99, 0xf9, 0x0b, 0x00, 0x3e, 0x64, 0x8c, 0x61,
	0x00, 0x7f, 0x63, 0xe6, 0x97, 0x01, 0xfc, 0x62, 0xef, 0xde, 0xbd, 0xcf,
	0x1f, 0x38, 0x70, 0xa0, 0x98, 0xee, 0x2a, 0xe5, 0x68, 0x00, 0xde, 0x21,
	0xcb, 0xb2, 0x23, 0xd6, 0xda, 0x1f, 0x18, 0x63, 0x1a, 0x57, 0x79, 0xd9,
	0x99, 0xa2, 0x28, 0xbe, 0xbe, 0xb2, 0xb2, 0x72, 0xd6, 0xdb, 0xc2, 0x1c,
	0xd2, 0x00, 0x5c, 0x91, 0xe7, 0xf9, 0x63, 0xcc, 0xfc, 0xab, 0xed, 0xbc,
	0xd6, 0x5a, 0x3b, 0x34, 0xc6, 0x3c, 0xd2, 0x6e, 0xb7, 0x7f, 0xe7, 0x7a,
	0x5d, 0xae, 0x99, 0x69, 0x2f, 0x60, 0x16, 0x6c, 0x6c, 0x6c, 0x34, 0x98,
	0xf9, 0x87, 0xdb, 0x7d, 0xbd, 0x31, 0x66, 0xa7, 0xb5, 0xf6, 0xd7, 0x27,
	0x4e, 0x9c, 0xb8, 0xd7, 0xe5, 0xba, 0x7c, 0xd0, 0x00, 0x00, 0x38, 0x7f,
	0xfe, 0xfc, 0x7e, 0x00, 0xb7, 0x56, 0x19, 0x63, 0x8c, 0xd9, 0x0d, 0xe0,
	0x74, 0x9a, 0xa6, 0x89, 0x93, 0x45, 0x79, 0xa2, 0x01, 0x00, 0x60, 0xad,
	0xfd, 0x52, 0x9d, 0x71, 0xc6, 0x98, 0x8f, 0x01, 0xf8, 0xae, 0xf0, 0x72,
	0xbc, 0xd2, 0x00, 0x00, 0x60, 0xe6, 0xfb, 0xea, 0x8e, 0x25, 0xa2, 0x95,
	0x63, 0xc7, 0x8e, 0x2d, 0x48, 0xae, 0xc7, 0x27, 0x0d, 0x00, 0x00, 0x63,
	0xcc, 0xe2, 0x04, 0xc3, 0xaf, 0xdf, 0xb5, 0x6b, 0xd7, 0x37, 0xc5, 0x16,
	0xe3, 0x59, 0xf4, 0x01, 0xc8, 0xf3, 0xfc, 0x26, 0x00, 0x13, 0xfd, 0x06,
	0x17, 0x45, 0xf1, 0x84, 0xd0, 0x72, 0xbc, 0x8b, 0x3e, 0x00, 0xd6, 0xda,
	0x0f, 0x4c, 0x5a, 0xc3, 0x18, 0x73, 0x67, 0x96, 0x65, 0xb7, 0x49, 0xac,
	0xc7, 0x37, 0x0d, 0x80, 0xb5, 0xef, 0x13, 0x2a, 0xf5, 0x88, 0x50, 0x1d,
	0xaf, 0xa2, 0x0f, 0x80, 0x31, 0xe6, 0xbd, 0x12, 0x75, 0x98, 0xf9, 0xf3,
	0x12, 0x75, 0x7c, 0x8b, 0x3e, 0x00, 0x00, 0xae, 0x97, 0x28, 0xc2, 0xcc,
	0xf7, 0x48, 0xd4, 0xf1, 0x2d, 0xfa, 0x00, 0x10, 0x91, 0xc8, 0x25, 0x9c,
	0x31, 0xe6, 0xb6, 0x5e, 0xaf, 0xb7, 0x47, 0xa2, 0x96, 0x4f, 0xd1, 0x07,
	0x80, 0x99, 0xc5, 0xae, 0xe1, 0x9b, 0xcd, 0xe6, 0x47, 0xa4, 0x6a, 0xf9,
	0x12, 0x7d, 0x00, 0x88, 0xe8, 0x06, 0xc1, 0x5a, 0xb7, 0x48, 0xd5, 0xf2,
	0x25, 0xfa, 0x00, 0x00, 0xb8, 0x51, 0xaa, 0x90, 0xb5, 0x56, 0x03, 0x10,
	0x1a, 0x66, 0x16, 0xdb, 0xb7, 0x99, 0xf9, 0xfd, 0x52, 0xb5, 0x7c, 0x89,
	0x3e, 0x00, 0x44, 0xb4, 0x57, 0xaa, 0x96, 0x31, 0x46, 0x6c, 0x3b, 0xf1,
	0x25, 0xfa, 0x00, 0x58, 0x6b, 0x3f, 0x28, 0x55, 0x8b, 0x99, 0x35, 0x00,
	0x21, 0xd9, 0xd8, 0xd8, 0x68, 0x18, 0x63, 0x6e, 0x95, 0xaa, 0xc7, 0xcc,
	0xbb, 0xa5, 0x6a, 0xf9, 0x12, 0x75, 0x00, 0xce, 0x9d, 0x3b, 0x77, 0x33,
	0x80, 0x1d, 0x82, 0x25, 0x83, 0xfb, 0x79, 0x06, 0xb7, 0x60, 0x49, 0x44,
	0x74, 0xe7, 0xb4, 0xd7, 0x30, 0x6d, 0x51, 0x07, 0xc0, 0x5a, 0xfb, 0x19,
	0xc9, 0x7a, 0xc6, 0x98, 0xe0, 0x1e, 0xb2, 0x8d, 0x3a, 0x00, 0x00, 0x3e,
	0x2b, 0x5c, 0x6f, 0x2c, 0x5c, 0xcf, 0xb9, 0x68, 0x03, 0xd0, 0xeb, 0xf5,
	0xf6, 0x10, 0x91, 0x74, 0x07, 0xef, 0xdf, 0xc2, 0xf5, 0x9c, 0x8b, 0x36,
	0x00, 0xcd, 0x66, 0xf3, 0x61, 0x00, 0xef, 0x11, 0x2e, 0xfb, 0x2f, 0xe1,
	0x7a, 0xce, 0x45, 0x1b, 0x00, 0x00, 0xdf, 0x92, 0x2e, 0xc8, 0xcc, 0x7a,
	0x06, 0x08, 0x41, 0x9e, 0xe7, 0xf7, 0x00, 0x10, 0x7f, 0x80, 0x83, 0x88,
	0x2e, 0x4a, 0xd7, 0x74, 0x2d, 0xca, 0x00, 0x58, 0x6b, 0xbf, 0xe7, 0xa2,
	0x2e, 0x33, 0x9f, 0x77, 0x51, 0xd7, 0x25, 0xe9, 0x3d, 0x70, 0xe6, 0x65,
	0x59, 0xb6, 0x1f, 0xc0, 0xa3, 0x8e, 0xca, 0x07, 0x17, 0x80, 0xa8, 0xce,
	0x00, 0x83, 0xc1, 0x60, 0x37, 0x80, 0xdc, 0xe1, 0x14, 0x1a, 0x80, 0x59,
	0x76, 0xf9, 0xf2, 0xe5, 0x1c, 0xc0, 0xed, 0xae, 0xea, 0x1b, 0x63, 0xfe,
	0xe2, 0xaa, 0xb6, 0x2b, 0xc1, 0xdd, 0xb9, 0xaa, 0x2b, 0xcf, 0xf3, 0x25,
	0x66, 0xfe, 0xb9, 0xc3, 0x29, 0x46, 0x87, 0x0f, 0x1f, 0x6e, 0x11, 0x11,
	0x3b, 0x9c, 0x43, 0x5c, 0x14, 0x67, 0x80, 0x2c, 0xcb, 0xf6, 0x15, 0x45,
	0xb1, 0xe6, 0x72, 0x0e, 0x6b, 0xed, 0x9b, 0xa1, 0x1d, 0x7c, 0x20, 0x82,
	0x00, 0x64, 0x59, 0xf6, 0x39, 0x00, 0xcf, 0x18, 0x63, 0x9a, 0x2e, 0xe7,
	0x21, 0xa2, 0x37, 0x5c, 0xd6, 0x77, 0x65, 0xae, 0x03, 0x90, 0xe7, 0xf9,
	0x17, 0x01, 0xfc, 0x16, 0x13, 0xfe, 0xed, 0xdf, 0x76, 0x30, 0xf3, 0xeb,
	0xae, 0xe7, 0x70, 0x61, 0x6e, 0x03, 0x90, 0x65, 0xd9, 0xe3, 0x45, 0x51,
	0x3c, 0x07, 0x0f, 0x07, 0x1f, 0x00, 0x8c, 0x31, 0x7f, 0xf6, 0x31, 0x8f,
	0xb4, 0xb9, 0xbb, 0x0f, 0xb0, 0xbe, 0xbe, 0xbe, 0x73, 0x6b, 0x6b, 0x2b,
	0x05, 0x70, 0xd0, 0x18, 0xaf, 0xf9, 0xd6, 0x00, 0x4c, 0x5b, 0x9a, 0xa6,
	0xf7, 0x6d, 0x6d, 0x6d, 0xfd, 0x0c, 0xc0, 0x1d, 0xbe, 0xe7, 0x1e, 0x8f,
	0xc7, 0xaf, 0xf9, 0x9e, 0x53, 0xc2, 0x5c, 0x5c, 0x06, 0xe6, 0x79, 0x7e,
	0x13, 0x33, 0x7f, 0x1f, 0xc0, 0x32, 0xa6, 0xb0, 0xad, 0x31, 0xf3, 0x5f,
	0x3b, 0x9d, 0xce, 0xcd, 0xbe, 0xe7, 0x95, 0x10, 0xf4, 0x19, 0x60, 0x6d,
	0x6d, 0xed, 0x86, 0xd1, 0x68, 0xf4, 0x9d, 0xa2, 0x28, 0x56, 0xa6, 0xfc,
	0x48, 0xf6, 0x99, 0x29, 0xce, 0x3d, 0x91, 0x20, 0x03, 0xd0, 0xef, 0xf7,
	0x17, 0x1b, 0x8d, 0xc6, 0xb7, 0x47, 0xa3, 0xd1, 0x13, 0x00, 0x16, 0x3c,
	0xef, 0xf5, 0xff, 0x83, 0x88, 0x34, 0x00, 0xae, 0xa5, 0x69, 0x7a, 0xbb,
	0x31, 0xe6, 0x61, 0x6b, 0xed, 0x57, 0x88, 0xe8, 0xfe, 0x69, 0xaf, 0xa7,
	0x24, 0xd8, 0x00, 0x4c, 0xfc, 0x1e, 0x20, 0x4d, 0xd3, 0xbb, 0x89, 0xe8,
	0x20, 0x80, 0x07, 0x00, 0xdc, 0x82, 0x09, 0x1f, 0xb3, 0x6e, 0xb7, 0xdb,
	0xef, 0x5a, 0x53, 0x96, 0x65, 0xc1, 0xdd, 0x5d, 0xf3, 0x64, 0x04, 0xe0,
	0x2c, 0x80, 0x5f, 0x8e, 0xc7, 0xe3, 0x41, 0xb7, 0xdb, 0xbd, 0x54, 0xa7,
	0x48, 0xed, 0x33, 0xc0, 0xea, 0xea, 0xea, 0x75, 0x3b, 0x76, 0xec, 0xc8,
	0x89, 0x68, 0xa9, 0x6e, 0x0d, 0x35, 0x91, 0x04, 0xc0, 0x5d, 0x00, 0xee,
	0x6a, 0x36, 0x9b, 0x4b, 0x83, 0xc1, 0x60, 0xff, 0xa1, 0x43, 0x87, 0xde,
	0xac, 0x5a, 0xa4, 0xd6, 0xe6, 0xb9, 0xba, 0xba, 0x7a, 0x5d, 0x92, 0x24,
	0xcf, 0xe9, 0xc1, 0x9f, 0x19, 0x1f, 0xb7, 0xd6, 0x3e, 0xbb, 0xbe, 0xbe,
	0xbe, 0xb3, 0xea, 0xc0, 0x5a, 0x01, 0x48, 0x92, 0xe4, 0x24, 0xe4, 0x1f,
	0xa9, 0x56, 0x93, 0xb9, 0x63, 0x73, 0x73, 0xf3, 0xc9, 0xaa, 0x83, 0x2a,
	0x07, 0xe0, 0xca, 0xf3, 0x74, 0xdf, 0xa8, 0x3a, 0x4e, 0xb9, 0x47, 0x44,
	0x5f, 0xad, 0x3a, 0xa6, 0x72, 0x00, 0xac, 0xb5, 0x07, 0xab, 0x8e, 0x51,
	0xde, 0x54, 0xbe, 0x03, 0x5a, 0x67, 0x0b, 0x78, 0xa0, 0xc6, 0x18, 0x35,
	0xa3, 0x2a, 0x07, 0x80, 0x99, 0x83, 0xfb, 0x37, 0x28, 0xb1, 0xb0, 0xd6,
	0x56, 0xee, 0x47, 0x54, 0x0e, 0x80, 0xeb, 0x07, 0x2b, 0x54, 0x7d, 0x44,
	0xf4, 0x74, 0xd5, 0x31, 0x73, 0xfb, 0x3c, 0x40, 0x6c, 0xac, 0xb5, 0xaf,
	0x02, 0x38, 0x55, 0x75, 0x9c, 0x06, 0x60, 0x0e, 0x58, 0x6b, 0x5f, 0x25,
	0xa2, 0x87, 0x3a, 0x9d, 0xce, 0xa8, 0xea, 0xd8, 0x60, 0x7a, 0x01, 0xea,
	0xdd, 0xac, 0xb5, 0x43, 0x00, 0x7f, 0x22, 0xa2, 0xa7, 0x89, 0xe8, 0x54,
	0x9d, 0x83, 0x0f, 0x68, 0x00, 0xa6, 0xae, 0xdc, 0xfb, 0xf0, 0x4d, 0xb7,
	0x80, 0xc8, 0x69, 0x00, 0x22, 0xa7, 0x01, 0x88, 0x9c, 0xbe, 0x07, 0x98,
	0xb2, 0x0a, 0xcf, 0x3b, 0x88, 0xf4, 0xff, 0xcb, 0xf4, 0x0c, 0x10, 0x8e,
	0xb7, 0xfb, 0xff, 0xc7, 0x9b, 0xcd, 0xe6, 0x99, 0xc1, 0x60, 0x20, 0x72,
	0x47, 0x56, 0x03, 0x10, 0xa6, 0xda, 0xfd, 0xff, 0x32, 0x0d, 0x40, 0xb8,
	0x6a, 0xf5, 0xff, 0xcb, 0x34, 0x00, 0x01, 0xab, 0xd3, 0xff, 0x2f, 0xd3,
	0x00, 0x84, 0x6d, 0xe2, 0xbf, 0x80, 0xd2, 0x00, 0x44, 0x4e, 0x03, 0x10,
	0xb0, 0x3a, 0xfd, 0xff, 0x32, 0x0d, 0x40, 0xc0, 0xea, 0xf4, 0xff, 0xcb,
	0x34, 0x00, 0x81, 0xaa, 0xdb, 0xff, 0x2f, 0xd3, 0x00, 0x04, 0x68, 0x92,
	0xfe, 0x7f, 0x99, 0xde, 0x0a, 0x0e, 0x84, 0x54, 0xff, 0xbf, 0x4c, 0x03,
	0xe0, 0xd9, 0xb4, 0xfb, 0xff, 0x65, 0xba, 0x05, 0x44, 0x4e, 0x03, 0x10,
	0x39, 0x0d, 0x40, 0xe4, 0xf4, 0x3d, 0x80, 0x67, 0xd3, 0xee, 0xff, 0x97,
	0xe9, 0x19, 0x60, 0x76, 0x39, 0xe9, 0xff, 0x97, 0x69, 0x00, 0xc2, 0x20,
	0xd6, 0xff, 0x2f, 0xd3, 0x00, 0x84, 0x43, 0xa4, 0xff, 0x5f, 0xa6, 0x01,
	0x08, 0x88, 0x44, 0xff, 0xbf, 0x4c, 0x03, 0x10, 0x16, 0xf1, 0xff, 0x80,
	0xaa, 0x01, 0x88, 0x9c, 0x06, 0x20, 0x20, 0x12, 0xfd, 0xff, 0x32, 0x0d,
	0x40, 0x40, 0x24, 0xfa, 0xff, 0x65, 0x1a, 0x80, 0x40, 0x48, 0xf5, 0xff,
	0xcb, 0x34, 0x00, 0x01, 0x90, 0xec, 0xff, 0x97, 0xe9, 0xad, 0xe0, 0x19,
	0xe5, 0xaa, 0xff, 0x5f, 0xe6, 0x2a, 0x00, 0x67, 0x89, 0xe8, 0xc8, 0x70,
	0x38, 0x7c, 0xa1, 0xd5, 0x6a, 0x11, 0x33, 0xef, 0x63, 0xe6, 0xe3, 0x00,
	0x16, 0x1d, 0xcd, 0x37, 0xb3, 0x66, 0xad, 0xff, 0x5f, 0xe6, 0x22, 0x00,
	0x67, 0x93, 0x24, 0xb9, 0x7f, 0x79, 0x79, 0xf9, 0x9f, 0xef, 0xf8, 0xda,
	0x6f, 0xfa, 0xfd, 0xfe, 0x8b, 0x8d, 0x46, 0xe3, 0x25, 0x44, 0x18, 0x82,
	0x59, 0x26, 0xfe, 0x1e, 0x80, 0x88, 0x8e, 0x94, 0x0e, 0x3e, 0x00, 0x60,
	0x65, 0x65, 0xe5, 0x1f, 0xcc, 0xfc, 0x94, 0xf4, 0x7c, 0x6a, 0x32, 0xe2,
	0x01, 0x18, 0x0e, 0x87, 0x2f, 0x5c, 0xe5, 0x7b, 0xbf, 0x97, 0x9e, 0x4f,
	0x4d, 0x46, 0x7c, 0x0b, 0x68, 0xb5, 0x5a, 0x33, 0xbd, 0xe7, 0xf9, 0x76,
	0x95, 0xfe, 0xff, 0x88, 0x99, 0x5f, 0x23, 0xa2, 0xd3, 0x0b, 0x0b, 0x0b,
	0xa7, 0x96, 0x96, 0x96, 0x86, 0x5e, 0x17, 0x76, 0x85, 0xf8, 0x19, 0x80,
	0x99, 0xf7, 0xfd, 0xbf, 0xef, 0x25, 0x49, 0xf2, 0xa0, 0xf4, 0x7c, 0x01,
	0x4b, 0x88, 0xe8, 0xd3, 0x00, 0x7e, 0xb4, 0xb9, 0xb9, 0xf9, 0xc7, 0x93,
	0x27, 0x4f, 0x7e, 0x78, 0x1a, 0x8b, 0x70, 0x11, 0x80, 0xe3, 0xfd, 0x7e,
	0xff, 0xc6, 0xf2, 0xd7, 0x7b, 0xbd, 0xde, 0x1e, 0x63, 0xcc, 0x31, 0xe9,
	0xf9, 0xe6, 0x01, 0x11, 0x7d, 0xa2, 0x28, 0x8a, 0x67, 0xd3, 0x34, 0x4d,
	0x7c, 0xcf, 0xed, 0xe2, 0x2a, 0x60, 0xb1, 0xd1, 0x68, 0xbc, 0x94, 0xa6,
	0xe9, 0x53, 0x6f, 0xef, 0xf9, 0x49, 0x92, 0x3c, 0x78, 0xe5, 0xe0, 0x7f,
	0xd4, 0xc1, 0x7c, 0xf3, 0xe2, 0x93, 0x44, 0xf4, 0x24, 0x80, 0xbe, 0xcf,
	0x49, 0xeb, 0x04, 0xe0, 0x2d, 0x5c, 0xfb, 0x73, 0x81, 0x16, 0x89, 0xe8,
	0x99, 0x56, 0xab, 0x55, 0xa3, 0x7c, 0xd4, 0xbe, 0x06, 0xcf, 0x01, 0xa8,
	0xb3, 0x05, 0x54, 0xfe, 0x5c, 0x1a, 0xb5, 0x6d, 0xde, 0x3f, 0xf1, 0xb4,
	0xce, 0x07, 0x46, 0xe8, 0xa5, 0xdc, 0x1c, 0xa9, 0x1c, 0x00, 0x22, 0x5a,
	0x73, 0xb1, 0x10, 0xe5, 0xa6, 0xdf, 0x7f, 0x2d, 0x95, 0x03, 0xd0, 0xe9,
	0x74, 0x5e, 0x06, 0xf0, 0x53, 0x07, 0x6b, 0x89, 0x5e, 0xa3, 0xd1, 0x38,
	0xed, 0x7b, 0xce, 0x5a, 0x97, 0x81, 0xe3, 0xf1, 0xb8, 0x0d, 0xe0, 0x45,
	0xe1, 0xb5, 0xc4, 0xee, 0x95, 0x0b, 0x17, 0x2e, 0x88, 0xf7, 0xfb, 0xaf,
	0xa5, 0x56, 0x00, 0xba, 0xdd, 0xee, 0xa5, 0xf1, 0x78, 0xfc, 0x65, 0x00,
	0x3f, 0xb1, 0xd6, 0xea, 0x27, 0x7b, 0x4e, 0xee, 0x95, 0xf1, 0x78, 0xfc,
	0xd0, 0xd1, 0xa3, 0x47, 0xdf, 0xf2, 0x3d, 0x71, 0xed, 0x1b, 0x41, 0xdd,
	0x6e, 0xf7, 0x52, 0xbb, 0xdd, 0x5e, 0x36, 0xc6, 0x7c, 0x8a, 0x99, 0x7f,
	0x0c, 0xe0, 0x75, 0xfc, 0xf7, 0x12, 0x51, 0x6d, 0xcf, 0x25, 0x00, 0x7f,
	0x00, 0xd0, 0xb9, 0x78, 0xf1, 0xe2, 0xbd, 0xdd, 0x6e, 0xf7, 0xef, 0xd3,
	0x5e, 0x90, 0x52, 0x4a, 0x29, 0xa5, 0x94, 0x52, 0x4a, 0x29, 0xa5, 0x94,
	0x52, 0x4a, 0x29, 0xa5, 0xe6, 0xcf, 0x7f, 0x00, 0x31, 0xeb, 0x1b, 0xd7,
	0xd8, 0x49, 0x66, 0xe1, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44,
	0xae, 0x42, 0x60, 0x82,
}
