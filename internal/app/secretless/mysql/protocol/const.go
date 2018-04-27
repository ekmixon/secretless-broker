/*
MIT License

Copyright (c) 2017 Aleksandr Fedotov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package protocol

const (
	responseEof         = 0xfe
	responseOk          = 0x00
	responsePrepareOk   = 0x00
	ResponseErr         = 0xff
	responseLocalinfile = 0xfb

	// MySQL field types constants
	fieldTypeString   = 0xfd
	fieldTypeLongLong = 0x08
	fieldTypeDouble   = 0x05

	// There is no code for Resultset in MySQL internal protocol
	// so it's defined here for convenience
	responseResultset = 0xbb

	// MySQL connection state constants
	ConnStateStarted  = 0xf4
	ConnStateFinished = 0xf5

	// Digits after comma
	doubleDecodePrecision = 6
)

const (
	comQuit byte = iota + 1
	comInitDB
	ComQuery
	ComFieldList
	comCreateDB
	comDropDB
	comRefresh
	comShutdown
	comStatistics
	comProcessInfo
	comConnect
	comProcessKill
	comDebug
	comPing
	comTime
	comDelayedInsert
	comChangeUser
	comBinlogDump
	comTableDump
	comConnectOut
	comRegisterSlave
	ComStmtPrepare
	ComStmtExecute
	comStmtSendLongData
	ComStmtClose
	comStmtReset
	comSetOption
	comStmtFetch
)

// Capability flags
const (
	clientLongPassword uint32 = 1 << iota
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConnection
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)
