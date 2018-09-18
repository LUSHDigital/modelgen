package sqlfmt

var acronyms = [...]string{
	"amd",     // Advanced Micro Devices
	"api",     // application programming interface
	"arp",     // address resolution protocol
	"arpanet", // Advanced Research Projects Agency Network
	"as",      // autonomous system
	"ascii",   // American Standard Code for Information Interchange
	"att",     // American Telephone and Telegraph Company
	"ata",     // advanced technology attachment
	"atm",     // asynchronous transfer mode

	"b",     // byte
	"belug", // Bellevue Linux Users Group
	"bgp",   // border gateway protocol
	"bios",  // basic input output system
	"bkl",   // big kernel lock
	"bnc",   // Bayonet Neill-Concelman
	"bsa",   // Business Software Alliance
	"bsd",   // Berkeley Software Distribution (originally Berkeley Source Distribution)

	"ccitt",  // Comité Consultatif International Téléphonique et Télégraphique
	"cd",     // change directory | compact disc
	"cdn",    // content delivery network
	"cdrom",  // compact disc read-only memory
	"cjkv",   // Chinese Japanese Korean Vietnamese
	"cldr",   // common locale data repository
	"cli",    // command line interface
	"cpu",    // central processing unit
	"crc",    // cyclic redundancy check
	"crt",    // cathode ray tube
	"csmacd", // carrier sense multiple access/collision detection
	"css",    // cascading style sheets
	"cups",   // common UNIX printing system
	"cvs",    // concurrent versions system

	"daemon", // disk and execution monitor*
	"dec",    // Digital Equipment Corporation
	"dhcp",   // dynamic host configuration protocol
	"dlc",    // data link control
	"dll",    // dynamic link library
	"dmca",   // Digital Millennium Copyright Act
	"dns",    // domain name service
	"dos",    // disk operating system
	"dram",   // dynamic random access memory
	"dvd",    // digital versatile disc (originally digital video disc)

	"edge",   // enhanced data GSM environment
	"eeprom", // electrically erasable read-only memory
	"eff",    // Electronic Frontier Foundation
	"egp",    // exterior gateway protocol
	"eula",   // end user license agreement

	"faq",  // frequently asked questions
	"fdd",  // floppy disk drive
	"fddi", // fiber distributed data interface
	"foss", // free open source software
	"fqdn", // fully qualified domain name
	"fs",   // filesystem
	"fsf",  // Free Software Foundation
	"ftp",  // file transfer protocol

	"gb",    // gigabit
	"gb",    // gigabyte
	"gbe",   // gigabit Ethernet
	"gcc",   // GNU Compiler Collection (originally GNU C Compiler)
	"gfdl",  // GNU Free Documentation License
	"gid",   // group identification
	"gif",   // graphics interchange format
	"gigo",  // garbage in garbage out
	"gimp",  // GNU Image Manipulation Program
	"gnome", // GNU Network Object Model Environment
	"gocc",  // Government Open Code Collaborative
	"gnu",   // GNU's Not UNIX
	"gpg",   // GNU Privacy Guard
	"gprs",  // general packet radio service
	"grub",  // grand unified bootloader
	"gui",   // graphical user interface
	"guid",  // globally unique identifier

	"hdd",  // hard disk drive
	"hdlc", // high level data link control
	"hfs",  // hierarchical file system
	"hp",   // Hewlett-Packard
	"html", // hypertext markup language
	"http", // hypertext transfer protocol

	"iana",  // Internet Assigned Numbers Authority
	"ibm",   // International Business Machines
	"ic",    // integrated circuit
	"icann", // Internet Corporation for Assigned Names and Numbers
	"icmp",  // Internet control message protocol
	"id",    // Identifier
	"ide",   // Integrated drive electronics
	"ieee",  // Institute of Electrical and Electronic Engineers
	"ietf",  // Internet Engineering Task Force
	"igp",   // Interior gateway protocol
	"igrp",  // Interior gateway routing protocol
	"imap",  // Internet message access protocol
	"io",    // Input/output
	"ip",    // Internet protocol
	"ipc",   // Inter-process communication
	"ipx",   // Internetwork packet exchange
	"irc",   // Internet relay chat
	"iso",   // International Organization for Standardization
	"isp",   // Internet service provider
	"it",    // Information technology
	"ixp",   // Internet exchange points

	"jpeg", // Joint Photographic Experts Group
	"jvm",  // Java virtual machine
	"json", // JavaScript Object Notation

	"kb",  // kilobit
	"kb",  // kilobyte
	"kcp", // kernel control path
	"kde", // K Desktop Environment

	"lamp",   // Linux Apache MySQL and PHP
	"lan",    // local area network
	"lanana", // Linux Assigned Names and Numbers Authority
	"lap",    // link access procedure
	"lcd",    // liquid crystal display
	"ldap",   // lightweight directory access protocol
	"led",    // light emitting diode
	"lfs",    // Linux From Scratch
	"lgpl",   // GNU Lesser General Public License
	"lids",   // Linux intrusion detection system
	"lilo",   // Linux loader
	"linfo",  // The Linux Information Project
	"lfs",    // log-structured file system
	"lsb",    // Linux Mark Institute
	"lsb",    // Linux Standards Base
	"lsi",    // large scale integrated circuit
	"lug",    // Linux users group
	"lvm",    // logical volume management
	"lzw",    // Lempel-Ziv-Welch

	"mac",    // media access control
	"mb",     // megabit
	"mb",     // megabyte
	"mbr",    // master boot record
	"md5",    // message digest 5
	"mdi",    // medium dependent interface
	"mit",    // Massachussets Institute of Technology
	"mmu",    // memory management unit
	"ms",     // Microsoft
	"ms-dos", // Microsoft disk operating system
	"mtu",    // maximum transmission unit

	"nfs",   // network file system
	"nic",   // network interface card
	"nilfs", // new implementation log-structured file system
	"nis",   // network information system
	"nntp",  // network news transfer protocol
	"ntp",   // network time protocol

	"odf",  // open document format
	"os",   // operating system
	"osdl", // Open Source Development Labs
	"osi",  // open systems interconnection
	"ospf", // open shortest path first

	"pam",    // pluggable authentication modules
	"pcmcia", // Personal Computer Memory Card International Association
	"pda",    // personal digital assistant
	"pdf",    // portable document format
	"pgp",    // pretty good privacy
	"php",    // PHP hypertext preprocessor (originally personal home page)
	"pid",    // process identification number
	"ping",   // packet Internet groper (orginally not an acronym)
	"pki",    // public key cryptography
	"pki",    // public key infrastructure
	"png",    // portable network graphics
	"pnp",    // plug-and-play
	"pop",    // post office protocol
	"posix",  // portable operating system interface
	"pots",   // plain old telephone service
	"ppp",    // point-to-point protocol
	"ps",     // postscript
	"pstn",   // public switched telephone network
	"pwd",    // print working directory

	"raid",  // redundant arrays of independent disks
	"ram",   // random access memory
	"rarp",  // reverse address resolution protocol
	"rdbms", // relational database management system
	"rfid",  // radio frequency identification
	"rhce",  // Red Hat Certified Engineer
	"rip",   // routing information protocol
	"rj",    // registered jack
	"rmon",  // remote monitoring
	"rms",   // Richard M. Stallman
	"rom",   // read-only memory
	"rpc",   // remote procedure call
	"rpm",   // Red Hat package manager
	"rss",   // really simple syndication
	"rtos",  // real time operating system
	"rtp",   // real-time transport protocol

	"san",   // storage area network
	"sane",  // scanner access now easy
	"sco",   // Santa Cruz Operation
	"scsi",  // small computer standard interface
	"sdlc",  // synchronous data link control
	"sdram", // synchronous dynamic random access memory
	"sgid",  // set group ID
	"sgml",  // standard generalized markup language
	"smb",   // server message block
	"snmp",  // simple network management protocol
	"smtp",  // simple mail transfer protocol
	"soap",  // simple object access protocol
	"spam",  // superfluous pieces of additional mail*
	"sram",  // static random access memory
	"sri",   // Stanford Research Institute
	"ssh",   // secure shell
	"ssl",   // secure sockets layer
	"su",    // substitute user
	"suid",  // set user ID
	"svid",  // System V interface definition

	"tar",    // tape archive
	"tb",     // terabyte
	"tcp",    // transmission control protocol
	"tcp/ip", // transmission control protocol/Internet protocol
	"tcl",    // tool command language
	"tco",    // total cost of ownership
	"tco",    // transmission control protocol
	"tlb",    // translation lookaside buffer
	"tld",    // top level domain
	"tron",   // The Real Time Operating System
	"tsl",    // transport layer security
	"ttl",    // time-to-live
	"ttl",    // transistor-transistor logic
	"tty",    // teletype terminal

	"ucb",   // University of California at Berkeley
	"ucita", // Uniform Computer Information Transactions Act
	"ucla",  // University of California at Los Angeles
	"ucs",   // universal character set
	"udp",   // user datagram protocol
	"uid",   // user identification
	"uri",   // uniform resource identifier
	"url",   // uniform resource locator
	"usb",   // universal serial bus
	"utf",   // UCS transformation format
	"utms",  // universal mobile telecommunications system
	"uucp",  // UNIX-to-UNIX copy
	"uuid",  // univerally unique identifier

	"vfat", // virtual file allocation table
	"vga",  // video graphics array
	"vlsi", // very large scale integrated circuit
	"vm",   // virtual memory
	"vpn",  // virtual private network

	"w3c",  // World Wide Web Consortium
	"wan",  // wide area network
	"wap",  // wireless access point
	"wap",  // wireless application protocol
	"wep",  // wired equivalent privacy
	"wine", // WINE is not an emulator
	"wlan", // wireless local area network
	"www",  // World Wide Web
	"wxga", // wide extended graphics array

	"x",   // X Window System
	"xml", // extensible markup language
}
