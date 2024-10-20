// This package contains variables and functions to assist with testing
package test

import (
	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

const (
	ExampleFeedMessageJSObject string = `{
		sequenceNumber: 241801449,
		message: {
			message: {
				header: {
					kind: 3,
					sender: '0xa4b000000000000000000073657175656e636572',
					blockNumber: 20969771,
					timestamp: 1728980884,
					requestId: null,
					baseFeeL1: null
				},
				l2Msg: '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'
			},
			delayedMessagesRead: 1719783
		}
	}`
	ExampleFeedMessageJSON string = `{
		"sequenceNumber":241801449,
		"message":{
			"message":{
				"header":{
					"kind":3,
					"sender":"0xa4b000000000000000000073657175656e636572",
					"blockNumber":20969771,
					"timestamp":1728980884,
					"requestId":null,
					"baseFeeL1":null
				},
				"l2Msg":"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"
			},
				"delayedMessagesRead":1719783
			}
		}`
	ExampleFeedMessageString string = `{"sequenceNumber":241801449,"message":{"message":{"header":{"kind":3,"sender":"0xa4b000000000000000000073657175656e636572","blockNumber":20969771,"timestamp":1728980884,"requestId":null,"baseFeeL1":null},"l2Msg":"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"},"delayedMessagesRead":1719783}}`
)

var ExampleParsedMessage feedtypes.IncomingMessage = feedtypes.IncomingMessage{
	SequenceNumber: 241801449,
	Message: feedtypes.MessageWrapper{
		Message: feedtypes.Message{
			Header: &arbostypes.L1IncomingMessageHeader{
				Kind:        3,
				Poster:      common.HexToAddress("0xA4b000000000000000000073657175656e636572"),
				BlockNumber: 20969771,
				Timestamp:   1728980884,
				RequestId:   nil,
				L1BaseFee:   nil,
			},
			L2msg: "AwAAAAAAAAQ2BPkEMg2DnSpggwa/HJQRlc9l+Ds6V2jzxJbToFrWQSxkt4ZENkxbsAC5A8TRI7TYAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADQAAAAAAAAAAAAAAAAAAAAAB7jppTjy9B3qWvPytoCv/j7WZNNC/l5VdcNkyeSZ/l6eCmHjCCSiT8crRk1eXJw2MAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEQ2TFuwAAAAAAAAAAAAAAAAAKa5gE6Sk4646XsR91cB3OEunsimAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAC4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAhNZXRhTWFzawAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA1aXBmczovL1FtZjg0QjFVY2hVTGllVGhUWDlxVXJNelFXMnNhaUM1YWhueldBYm1RWm1NVnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABB/WW+bTtWXAdY8joApgrS2RYKzmWMym7l/frWwF5B/cMfv/pRNS0QPJucH/qq2vAJh6zPOgouDbINFvxVaePSgRwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgwFJhaD72EZcmuijtNWiYL/Xn8J3Fy4Z2cF07LkWizIwZYsHSKB3IniQRufmfkTldsP8bIDxKbQsniULJSHWPl6CdF9p3QAAAAAAAAC0BAL4sIKksYKdXYCEATEtAIMHoSCUgq9JRH2KB+O9lb0NVvNSQVI/urGAuEQJXqezAAAAAAAAAAAAAAAA5ZJCegrs6S3j7e4fGOAVfAWGFWQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAfQ42qBgAMCAoLyewVn7q3WFrc0uWVxd1mVMgQ6+z0+eksgJhivHyPEtoH5Jz0TXyUqVphrVSqdSr6dpama3PiwRbvtkMckZGOSdAAAAAAAAAO4E+OsFg5ibU4MXBoyUZI/UqaAgDcMOMHrYDhIIyHgLN/iAuIQuEk6bAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGrJGR6AAAAAAAAAAAAAAAABpLP/tLtPV+/BkydZVb9Q9d3BUEQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADN+RoCDAUmGoBq8KBayS3dQtoaXtUG6nPFMIlXBd88H4lpQZEYoH1gvoHCmY8fGAwq+m7ymzVJbSJ/hWPq1JRtCXjNTQTbPyBwrAAAAAAAAAJEEAviNgqSxgyHxm4OYloCEATEtAIMPQk6UgRPJJR1Y/Rdn3mXXbxsMtzZfGDCAngAAAAH+pZkAAADKFs2ydLr4ABXSF8F/iXoAAAAAGsCAoL9jj2XOKxs42NtNCZMoE4Riee2hVjuSR+uxbb+H9w5MoE6fsP9YThaNeR/wMnLt44KgExP8hel+b3itwj+hjLEDAAAAAAAAALAE+K2CAX6Dxl1AgwNmSpSRLOWRRBkcEgTmRVn+glOg5J5lSIC4RAlep7MAAAAAAAAAAAAAAAAMjHe3/0wq9/bOu+ZzUKSQ491sswAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGxrk1uLvUAAAgwFJhqCQ8TWqtxP+LNo6mqRxXthv1CLzaqSlLOuwr85IDfqhKqAQefRyGtld45EgrYu7herUp/Z21reJxpCp7MJ2lF6XfQAAAAAAAAB7BAL4d4KksYMHdzOEdzWUAIRHhowAgwLNDpSh0tGOBViJsGrfkm48z0zKQwogwocCszdKB4AAgMABoGUpy3tXH2Bh0MYoXhkfoH/hGTJZ9DdrSqcwuUzdpvsLoC0XU8+aI/2nR2fSSzXd84MX39vMjiK06el93o1l6EE/",
		},
		DelayedMessagesRead: 1719783,
	},
	Signature: nil,
}

var ExampleDecodedL2Message []byte = []byte{
	3, 0, 0, 0, 0, 0, 0, 4, 54, 4, 249, 4, 50, 13, 131, 157, 42, 96, 131, 6,
	191, 28, 148, 17, 149, 207, 101, 248, 59, 58, 87, 104, 243, 196, 150, 211,
	160, 90, 214, 65, 44, 100, 183, 134, 68, 54, 76, 91, 176, 0, 185, 3, 196,
	209, 35, 180, 216, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 123, 142, 154, 83, 143, 47, 65, 222, 165,
	175, 63, 43, 104, 10, 255, 227, 237, 102, 77, 52, 47, 229, 229, 87, 92, 54,
	76, 158, 73, 159, 229, 233, 224, 166, 30, 48, 130, 74, 36, 252, 114, 180,
	100, 213, 229, 201, 195, 99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 68, 54, 76, 91, 176, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 166, 185, 128, 78, 146, 147, 142, 184, 233, 123, 17, 247, 87,
	1, 220, 225, 46, 158, 200, 166, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 64, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 2, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 224, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	8, 77, 101, 116, 97, 77, 97, 115, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 53, 105, 112, 102,
	115, 58, 47, 47, 81, 109, 102, 56, 52, 66, 49, 85, 99, 104, 85, 76, 105,
	101, 84, 104, 84, 88, 57, 113, 85, 114, 77, 122, 81, 87, 50, 115, 97, 105,
	67, 53, 97, 104, 110, 122, 87, 65, 98, 109, 81, 90, 109, 77, 86, 112, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 65, 253, 101, 190, 109, 59, 86, 92, 7, 88,
	242, 58, 0, 166, 10, 210, 217, 22, 10, 206, 101, 140, 202, 110, 229, 253,
	250, 214, 192, 94, 65, 253, 195, 31, 191, 250, 81, 53, 45, 16, 60, 155, 156,
	31, 250, 170, 218, 240, 9, 135, 172, 207, 58, 10, 46, 13, 178, 13, 22, 252,
	85, 105, 227, 210, 129, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 131, 1, 73, 133, 160, 251, 216,
	70, 92, 154, 232, 163, 180, 213, 162, 96, 191, 215, 159, 194, 119, 23, 46,
	25, 217, 193, 116, 236, 185, 22, 139, 50, 48, 101, 139, 7, 72, 160, 119, 34,
	120, 144, 70, 231, 230, 126, 68, 229, 118, 195, 252, 108, 128, 241, 41, 180,
	44, 158, 37, 11, 37, 33, 214, 62, 94, 130, 116, 95, 105, 221, 0, 0, 0, 0, 0,
	0, 0, 180, 4, 2, 248, 176, 130, 164, 177, 130, 157, 93, 128, 132, 1, 49, 45,
	0, 131, 7, 161, 32, 148, 130, 175, 73, 68, 125, 138, 7, 227, 189, 149, 189,
	13, 86, 243, 82, 65, 82, 63, 186, 177, 128, 184, 68, 9, 94, 167, 179, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 229, 146, 66, 122, 10, 236, 233, 45, 227, 237,
	238, 31, 24, 224, 21, 124, 5, 134, 21, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 244, 56, 218, 160, 96, 0, 192,
	128, 160, 188, 158, 193, 89, 251, 171, 117, 133, 173, 205, 46, 89, 92, 93,
	214, 101, 76, 129, 14, 190, 207, 79, 158, 146, 200, 9, 134, 43, 199, 200,
	241, 45, 160, 126, 73, 207, 68, 215, 201, 74, 149, 166, 26, 213, 74, 167,
	82, 175, 167, 105, 106, 102, 183, 62, 44, 17, 110, 251, 100, 49, 201, 25,
	24, 228, 157, 0, 0, 0, 0, 0, 0, 0, 238, 4, 248, 235, 5, 131, 152, 155, 83,
	131, 23, 6, 140, 148, 100, 143, 212, 169, 160, 32, 13, 195, 14, 48, 122,
	216, 14, 18, 8, 200, 120, 11, 55, 248, 128, 184, 132, 46, 18, 78, 155, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	26, 178, 70, 71, 160, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 105, 44, 255, 237,
	46, 211, 213, 251, 240, 100, 201, 214, 85, 111, 212, 61, 119, 112, 84, 17,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 51, 126, 70, 128, 131, 1, 73, 134, 160, 26,
	188, 40, 22, 178, 75, 119, 80, 182, 134, 151, 181, 65, 186, 156, 241, 76,
	34, 85, 193, 119, 207, 7, 226, 90, 80, 100, 70, 40, 31, 88, 47, 160, 112,
	166, 99, 199, 198, 3, 10, 190, 155, 188, 166, 205, 82, 91, 72, 159, 225, 88,
	250, 181, 37, 27, 66, 94, 51, 83, 65, 54, 207, 200, 28, 43, 0, 0, 0, 0, 0,
	0, 0, 145, 4, 2, 248, 141, 130, 164, 177, 131, 33, 241, 155, 131, 152, 150,
	128, 132, 1, 49, 45, 0, 131, 15, 66, 78, 148, 129, 19, 201, 37, 29, 88, 253,
	23, 103, 222, 101, 215, 111, 27, 12, 183, 54, 95, 24, 48, 128, 158, 0, 0, 0,
	1, 254, 165, 153, 0, 0, 0, 202, 22, 205, 178, 116, 186, 248, 0, 21, 210, 23,
	193, 127, 137, 122, 0, 0, 0, 0, 26, 192, 128, 160, 191, 99, 143, 101, 206,
	43, 27, 56, 216, 219, 77, 9, 147, 40, 19, 132, 98, 121, 237, 161, 86, 59,
	146, 71, 235, 177, 109, 191, 135, 247, 14, 76, 160, 78, 159, 176, 255, 88,
	78, 22, 141, 121, 31, 240, 50, 114, 237, 227, 130, 160, 19, 19, 252, 133,
	233, 126, 111, 120, 173, 194, 63, 161, 140, 177, 3, 0, 0, 0, 0, 0, 0, 0,
	176, 4, 248, 173, 130, 1, 126, 131, 198, 93, 64, 131, 3, 102, 74, 148, 145,
	44, 229, 145, 68, 25, 28, 18, 4, 230, 69, 89, 254, 130, 83, 160, 228, 158,
	101, 72, 128, 184, 68, 9, 94, 167, 179, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	12, 140, 119, 183, 255, 76, 42, 247, 246, 206, 187, 230, 115, 80, 164, 144,
	227, 221, 108, 179, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 27, 26, 228, 214, 226, 239, 80, 0, 0, 131, 1, 73, 134, 160, 144,
	241, 53, 170, 183, 19, 254, 44, 218, 58, 154, 164, 113, 94, 216, 111, 212,
	34, 243, 106, 164, 165, 44, 235, 176, 175, 206, 72, 13, 250, 161, 42, 160,
	16, 121, 244, 114, 26, 217, 93, 227, 145, 32, 173, 139, 187, 133, 234, 212,
	167, 246, 118, 214, 183, 137, 198, 144, 169, 236, 194, 118, 148, 94, 151,
	125, 0, 0, 0, 0, 0, 0, 0, 123, 4, 2, 248, 119, 130, 164, 177, 131, 7, 119,
	51, 132, 119, 53, 148, 0, 132, 71, 134, 140, 0, 131, 2, 205, 14, 148, 161,
	210, 209, 142, 5, 88, 137, 176, 106, 223, 146, 110, 60, 207, 76, 202, 67,
	10, 32, 194, 135, 2, 179, 55, 74, 7, 128, 0, 128, 192, 1, 160, 101, 41, 203,
	123, 87, 31, 96, 97, 208, 198, 40, 94, 25, 31, 160, 127, 225, 25, 50, 89,
	244, 55, 107, 74, 167, 48, 185, 76, 221, 166, 251, 11, 160, 45, 23, 83, 207,
	154, 35, 253, 167, 71, 103, 210, 75, 53, 221, 243, 131, 23, 223, 219, 204,
	142, 34, 180, 233, 233, 125, 222, 141, 101, 232, 65, 63,
}

// Transaction 1 (type: 0x0)
var ExampleTx1 *types.Transaction = types.NewTransaction(
	13, // Nonce
	common.HexToAddress("0x1195cf65f83b3a5768f3c496d3a05ad6412c64b7"), // To address
	big.NewInt(0x44364c5bb000), // Value (in Wei)
	0x6bf1c,                    // Gas limit
	big.NewInt(0x9d2a60),       // Gas price
	common.FromHex("0xd123b4d800000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000340000000000000000000000000000000007b8e9a538f2f41dea5af3f2b680affe3ed664d342fe5e5575c364c9e499fe5e9e0a61e30824a24fc72b464d5e5c9c363000000000000000000000000000000000000000000000000000044364c5bb000000000000000000000000000a6b9804e92938eb8e97b11f75701dce12e9ec8a60000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000002c000000000000000000000000000000000000000000000000000000000000002e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000084d6574614d61736b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000035697066733a2f2f516d6638344231556368554c696554685458397155724d7a515732736169433561686e7a5741626d515a6d4d567000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041fd65be6d3b565c0758f23a00a60ad2d9160ace658cca6ee5fdfad6c05e41fdc31fbffa51352d103c9b9c1ffaaadaf00987accf3a0a2e0db20d16fc5569e3d2811c00000000000000000000000000000000000000000000000000000000000000"), // Input data
)

// Transaction 2 (type: 0x2)
var ExampleToAddress2 common.Address = common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1")
var ExampleTx2 *types.Transaction = types.NewTx(&types.DynamicFeeTx{
	Nonce:    40285,
	GasTipCap: big.NewInt(0x0),
	GasFeeCap: big.NewInt(0x1312d00),
	Gas:      0x7a120,
	To:       &ExampleToAddress2,
	Value:    big.NewInt(0),
	Data:     common.FromHex("0x095ea7b3000000000000000000000000e592427a0aece92de3edee1f18e0157c058615640000000000000000000000000000000000000000000000000001f438daa06000"),
})

// Transaction 3 (type: 0x0)
var ExampleTx3 *types.Transaction = types.NewTransaction(
	5,
	common.HexToAddress("0x648fd4a9a0200dc30e307ad80e1208c8780b37f8"),
	big.NewInt(0x0),
	0x17068c,
	big.NewInt(0x989b53),
	common.FromHex("0x2e124e9b0000000000000000000000000000000000000000000000000000001ab24647a0000000000000000000000000692cffed2ed3d5fbf064c9d6556fd43d77705411000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000337e4680"),
)

// Transaction 4 (type: 0x2)
var ExampleToAddress4 common.Address = common.HexToAddress("0x8113c9251d58fd1767de65d76f1b0cb7365f1830")
var ExampleTx4 *types.Transaction = types.NewTx(&types.DynamicFeeTx{
	Nonce:    2224539,
	GasTipCap: big.NewInt(0x989680),
	GasFeeCap: big.NewInt(0x1312d00),
	Gas:      0xf424e,
	To:       &ExampleToAddress4,
	Value:    big.NewInt(0),
	Data:     common.FromHex("0x00000001fea599000000ca16cdb274baf80015d217c17f897a000000001a"),
})

// Transaction 5 (type: 0x0)
var ExampleTx5 *types.Transaction = types.NewTransaction(
	382,
	common.HexToAddress("0x912ce59144191c1204e64559fe8253a0e49e6548"),
	big.NewInt(0x0),
	0x3664a,
	big.NewInt(0xc65d40),
	common.FromHex("0x095ea7b30000000000000000000000000c8c77b7ff4c2af7f6cebbe67350a490e3dd6cb300000000000000000000000000000000000000000000001b1ae4d6e2ef500000"),
)

// Transaction 6 (type: 0x2)
var ExampleToAddress6 common.Address = common.HexToAddress("0xa1d2d18e055889b06adf926e3ccf4cca430a20c2")
var ExampleTx6 *types.Transaction = types.NewTx(&types.DynamicFeeTx{
	Nonce:    489267,
	GasTipCap: big.NewInt(0x77359400),
	GasFeeCap: big.NewInt(0x47868c00),
	Gas:      0x2cd0e,
	To:       &ExampleToAddress6,
	Value:    big.NewInt(0x2b3374a078000),
	Data:     []byte{}, // No input data
})

// Declare the full transactions slice
var ExampleTxns types.Transactions = types.Transactions{ExampleTx1, ExampleTx2, ExampleTx3, ExampleTx4, ExampleTx5, ExampleTx6}

var ExampleStartTx *types.Transaction = types.NewTx(&types.ArbitrumInternalTx{
	ChainId: big.NewInt(utils.ArbiturmChainId),
	Data: common.FromHex("0x6bf6a42d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013ff92b000000000000000000000000000000000000000000000000000000000fbc76320000000000000000000000000000000000000000000000000000000000000000"),
})

// TODO: v is defined as hard-coded way, v calculation with r, s, tx hash should be applied
func GetExampleSignedTxns() types.Transactions {
	chainId := big.NewInt(utils.ArbiturmChainId)
	legacySigner := types.NewEIP155Signer(chainId)
	londonSigner := types.NewLondonSigner(chainId)

	// ExampleTx1 (Legacy Tx)
	r := new(big.Int).SetBytes(common.Hex2Bytes("fbd8465c9ae8a3b4d5a260bfd79fc277172e19d9c174ecb9168b3230658b0748"))
	s := new(big.Int).SetBytes(common.Hex2Bytes("7722789046e7e67e44e576c3fc6c80f129b42c9e250b2521d63e5e82745f69dd"))
	v := big.NewInt(0)

	// Create signature for ExampleTx1
	sig := make([]byte, 65)
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64())
	signedTx1, _ := ExampleTx1.WithSignature(legacySigner, sig)

	// ExampleTx2 (EIP-1559 Dynamic Fee Tx)
	r.SetBytes(common.Hex2Bytes("bc9ec159fbab7585adcd2e595c5dd6654c810ebecf4f9e92c809862bc7c8f12d"))
	s.SetBytes(common.Hex2Bytes("7e49cf44d7c94a95a61ad54aa752afa7696a66b73e2c116efb6431c91918e49d"))
	v = big.NewInt(0)

	// Create signature for ExampleTx2
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64()) // For EIP-1559, this is yParity (0 or 1)
	signedTx2, _ := ExampleTx2.WithSignature(londonSigner, sig)

	// ExampleTx3 (Legacy Tx)
	r.SetBytes(common.Hex2Bytes("1abc2816b24b7750b68697b541ba9cf14c2255c177cf07e25a506446281f582f"))
	s.SetBytes(common.Hex2Bytes("70a663c7c6030abe9bbca6cd525b489fe158fab5251b425e33534136cfc81c2b"))
	v = big.NewInt(1)

	// Create signature for ExampleTx3
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64())
	signedTx3, _ := ExampleTx3.WithSignature(legacySigner, sig)

	// ExampleTx4 (EIP-1559 Dynamic Fee Tx)
	r.SetBytes(common.Hex2Bytes("bf638f65ce2b1b38d8db4d09932813846279eda1563b9247ebb16dbf87f70e4c"))
	s.SetBytes(common.Hex2Bytes("4e9fb0ff584e168d791ff03272ede382a01313fc85e97e6f78adc23fa18cb103"))
	v = big.NewInt(0)

	// Create signature for ExampleTx4
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64())
	signedTx4, _ := ExampleTx4.WithSignature(londonSigner, sig)

	// ExampleTx5 (Legacy Tx)
	r.SetBytes(common.Hex2Bytes("90f135aab713fe2cda3a9aa4715ed86fd422f36aa4a52cebb0afce480dfaa12a"))
	s.SetBytes(common.Hex2Bytes("1079f4721ad95de39120ad8bbb85ead4a7f676d6b789c690a9ecc276945e977d"))
	v = big.NewInt(1)

	// Create signature for ExampleTx5
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64())
	signedTx5, _ := ExampleTx5.WithSignature(legacySigner, sig)

	// ExampleTx6 (EIP-1559 Dynamic Fee Tx)
	r.SetBytes(common.Hex2Bytes("6529cb7b571f6061d0c6285e191fa07fe1193259f4376b4aa730b94cdda6fb0b"))
	s.SetBytes(common.Hex2Bytes("2d1753cf9a23fda74767d24b35ddf38317dfdbcc8e22b4e9e97dde8d65e8413f"))
	v = big.NewInt(1) // yParity is 1

	// Create signature for ExampleTx6
	r.FillBytes(sig[:32])
	s.FillBytes(sig[32:64])
	sig[64] = byte(v.Uint64())
	signedTx6, _ := ExampleTx6.WithSignature(londonSigner, sig)

	// Return signed transactions
	txns := types.Transactions{signedTx1, signedTx2, signedTx3, signedTx4, signedTx5, signedTx6}
	return txns
}

func CompareTransactionFields(tx1, tx2 *types.Transaction) bool {
	// Compare transaction types first
	if tx1.Type() != tx2.Type() {
		return false
	}

	// Compare the public fields, including transaction type-specific fields
	switch tx1.Type() {
	case types.LegacyTxType: // Legacy transactions (type 0x0)
		return tx1.Nonce() == tx2.Nonce() &&
			tx1.To().Hex() == tx2.To().Hex() &&
			tx1.Value().Cmp(tx2.Value()) == 0 &&
			tx1.Gas() == tx2.Gas() &&
			compareBigInts(tx1.GasPrice(), tx2.GasPrice()) &&
			bytes.Equal(tx1.Data(), tx2.Data()) &&
			compareSignatureValues(tx1, tx2) && // Compare v, r, s
			tx1.Hash() == tx2.Hash() &&         // Compare hash
			compareBigInts(tx1.ChainId(), tx2.ChainId()) // Compare chainId

	case types.DynamicFeeTxType: // EIP-1559 transactions (type 0x2)
		return tx1.Nonce() == tx2.Nonce() &&
			tx1.To().Hex() == tx2.To().Hex() &&
			tx1.Value().Cmp(tx2.Value()) == 0 &&
			tx1.Gas() == tx2.Gas() &&
			compareBigInts(tx1.GasTipCap(), tx2.GasTipCap()) &&
			compareBigInts(tx1.GasFeeCap(), tx2.GasFeeCap()) &&
			bytes.Equal(tx1.Data(), tx2.Data()) &&
			compareSignatureValues(tx1, tx2) && // Compare v, r, s
			tx1.Hash() == tx2.Hash() &&         // Compare hash
			compareBigInts(tx1.ChainId(), tx2.ChainId()) // Compare chainId

	default:
		// Return false for unsupported or unknown transaction types
		return false
	}
}

// Helper function to compare v, r, s values
func compareSignatureValues(tx1, tx2 *types.Transaction) bool {
	v1, r1, s1 := tx1.RawSignatureValues()
	v2, r2, s2 := tx2.RawSignatureValues()

	return compareBigInts(v1, v2) &&
		compareBigInts(r1, r2) &&
		compareBigInts(s1, s2)
}

// Helper function to compare *big.Int values
func compareBigInts(b1, b2 *big.Int) bool {
	if b1 == nil && b2 == nil {
		return true
	}
	if b1 == nil || b2 == nil {
		return false
	}
	return b1.Cmp(b2) == 0
}

// CompareArbitrumInternalTx compares two ArbitrumInternalTx transactions field by field.
func CompareArbitrumInternalTx(tx1, tx2 *types.Transaction) bool {
	// Check if both transactions have the type ArbitrumInternalTxType (0x6A)
	if tx1.Type() != types.ArbitrumInternalTxType || tx2.Type() != types.ArbitrumInternalTxType {
		// If either transaction is not of type ArbitrumInternalTxType, return false
		return false
	}

	// Compare ChainId fields
	if tx1.ChainId().Cmp(tx2.ChainId()) != 0 {
		return false
	}

	// Compare Data fields
	if !bytes.Equal(tx1.Data(), tx2.Data()) {
		return false
	}

	// If all fields match, return true
	return true
}

// Notice : Only v, r, s value is printed as fmt.Printf()
func PrintTransactionFields(tx1, tx2 *types.Transaction) bool {
	fmt.Println("Comparing transactions...")
	
	v1, r1, s1 := tx1.RawSignatureValues()
	v2, r2, s2 := tx2.RawSignatureValues()
	
	fmt.Printf("v1: %v, v2: %v\n", v1, v2)
	fmt.Printf("r1: %v, r2: %v\n", r1, r2)
	fmt.Printf("s1: %v, s2: %v\n", s1, s2)
	
	// Add your previous comparison code here
	
	return compareBigInts(v1, v2) &&
		compareBigInts(r1, r2) &&
		compareBigInts(s1, s2) &&
		tx1.Hash() == tx2.Hash() && // Compare hash
		compareBigInts(tx1.ChainId(), tx2.ChainId()) // Compare chainId
}
