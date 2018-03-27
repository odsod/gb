# Memory

## Memory Map

### ROM Bank 00

| Address     | Type               | Description                |
|-------------|--------------------|----------------------------|
| `0000-0007` | RST handlers       | RST 00                     |
| `0008-000f` |                    | RST 08                     |
| `0010-0017` |                    | RST 10                     |
| `0018-001f` |                    | RST 18                     |
| `0020-0027` |                    | RST 20                     |
| `0028-002f` |                    | RST 28                     |
| `0030-0037` |                    | RST 30                     |
| `0038-003f` |                    | RST 38                     |
| `0040-0047` | Interrupt handlers | VBlank                     |
| `0048-004f` |                    | LCDC Status                |
| `0050-0057` |                    | Timer Overflow             |
| `0058-005f` |                    | Serial Transfer Completion |
| `0060-0067` |                    | Joypad Input               |
| `0068-00ff` | Unused             |                            |
| `0100-0103` | ROM Header         | Code entry point           |
| `0104-0133` |                    | Nintendo logo              |
| `0134-0142` |                    | Game title                 |
| `0143`      |                    | Game Boy Color flag        |
| `0144-0145` |                    | New licensee code          |
| `0146`      |                    | Super Game Boy flag        |
| `0147`      |                    | Cartridge type             |
| `0148`      |                    | ROM Size                   |
| `0149`      |                    | RAM Size                   |
| `014a`      |                    | Destination code           |
| `014b`      |                    | Licensee code              |
| `014c`      |                    | Mark ROM Version number    |
| `014d`      |                    | Complement check           |
| `014e-014f` |                    | Checksum                   |
| `0150-3fff` | ROM                | Cartridge ROM              |

### ROM Bank 01-FF

| Address     | Description   |
|-------------|---------------|
| `4000-7fff` | Cartridge ROM |

### VRAM

| Address     | Description                        |
|-------------|------------------------------------|
| `8000-8fff` | Tile Pattern Table (unsigned mode) |
| `8800-97ff` | Tile Pattern Table (signed mode)   |
| `9800-9bff` | BG Tile Map 1                      |
| `9c00-9fff` | BG Tile Map 2                      |

### RAM

| Address     |               |
|-------------|---------------|
| `a000-bfff` | Cartridge RAM |
| `c000-cfff` | Internal RAM  |
| `e000-fdff` | Echo RAM      |

### OAM

| Address     | Description           |
|-------------|-----------------------|
| `fe00-fe9f` | 40 4-byte OAM entries |
| `fea0-feff` | Unused                |

### I/O Registers

| Address     | Type            | Description                        |
|-------------|-----------------|------------------------------------|
| `ff00`      | Joypad          | Joypad (P1) (R/W)                  |
| `ff01`      | Serial transfer | Serial transfer data (SB) (R/W)    |
| `ff02`      |                 | Serial transfer control (SC) (R/W) |
| `ff03`      | Unused          |                                    |
| `ff04`      | Timer           | [DIV](#div)                        |
| `ff05`      |                 | [TIMA](#tima)                      |
| `ff06`      |                 | [TMA](#tma)                        |
| `ff07`      |                 | [TAC](#tac)                        |
| `ff08`      | Interrupts      | Interrupt Flag (IF) (R/W)          |
| `ff09`      | Unused          |                                    |
| `ff10-ff3f` | Sound           | TODO                               |
| `ff40`      | Video           | [LCDC](#lcdc)                      |
| `ff41`      |                 | [STAT](#stat)                      |
| `ff42`      |                 | [SCY](#scy)                        |
| `ff43`      |                 | [SCX](#scx)                        |
| `ff44`      |                 | [LY](#ly)                          |
| `ff45`      |                 | [LYC](#lyc)                        |
| `ff46`      |                 | [DMA](#dma)                        |
| `ff47`      |                 | [BGP](#bgp)                        |
| `ff48`      |                 | [OBP0](#obp0)                      |
| `ff49`      |                 | [OBP1](#obp1)                      |
| `ff4a`      |                 | Window Y Position (WY) (R/W)       |
| `ff4b`      |                 | Window X Position (WX) (R/W)       |

### HRAM

| Address     | Description |
|-------------|-------------|
| `ff4c-fffe` | HRAM        |

### Interrupt Enable Register

| Address | Description                 |
|---------|-----------------------------|
| `ffff`  | Interrupt Enable (IE) (R/W) |

## Registers

### DIV

* **Purpose**: Divider Register
* **R/W**: Read & Write
* **Address**: `ff04`

Increments at a rate of 16384Hz.

GameBoy clock speed is 4194304Hz, meaning it increments every 256th CPU
cycle.

Can be implemented by a 16-bit cycle counter, by shifting out the low 8
bits upon read.

Writing any value resets to 0.

### TIMA

* **Purpose**: Timer Counter
* **R/W**: Read & Write
* **Address**: `ff05`

Incremented at a frequency determined by [TAC](#tac).

Reset to [TMA](#tma) and triggers an interrupt upon overflow.

### TMA

* **Purpose**: Timer Modulo
* **R/W**: Read & Write
* **Address**: `ff06`

Loaded into [TIMA](#tima) when it overflows.

### TAC

* **Purpose**: Timer Control
* **R/W**: Read & Write
* **Address**: `ff07`

| Bits | Purpose            | When 0    | When 1    |
|------|--------------------|-----------|-----------|
| 2    | Timer Enable       | Off       | On        |
| 1-0  | Input Clock Select | See below | See below |

| Input Clock Select | Clock Frequency | [DIV](#div) bit |
|--------------------|-----------------|-----------------|
| 00                 | 4096Hz          | 9               |
| 01                 | 262144Hz        | 3               |
| 10                 | 65536Hz         | 5               |
| 11                 | 16384Hz         | 7               |

[TMA](#tma) is incremented on the falling edge of the selected [DIV](#div)
bit.

### LCDC

* **Purpose**: LCD Control
* **R/W**: Read & Write
* **Address**: `ff40`
* **Default**: `91`

| Bit | Function                      | When 0      | When 1      | Default |
|-----|-------------------------------|-------------|-------------|---------|
| 7   | LCD display                   | Off         | On          | `1`     |
| 6   | Window tile map               | `9800-9bff` | `9c00-9fff` | `0`     |
| 5   | Window display                | Off         | On          | `0`     |
| 4   | Background & window tile data | `8800-97ff` | `8000-8fff` | `1`     |
| 3   | Background tile map           | `9800-9bff` | `9c00-9fff` | `0`     |
| 2   | Sprite size                   | 8x8         | 8x16        | `0`     |
| 1   | Sprite display                | Off         | On          | `0`     |
| 0   | Background display            | Off         | On          | `1`     |

### STAT

* **Purpose**: LCD Status
* **R/W**: Read & Write
* **Address**: `ff41`

| Bit | Function                  | When 0                   | When 1                  |
|-----|---------------------------|--------------------------|-------------------------|
| 7   | Unused                    |                          |                         |
| 6   | LYC Coincidence Interrupt | Off                      | On                      |
| 5   | Mode 2 OAM Interrupt      | Off                      | On                      |
| 4   | Mode 1 V-Blank Interrupt  | Off                      | On                      |
| 3   | Mode 0 H-Blank Interrupt  | Off                      | On                      |
| 2   | Coincidence Flag          | [LYC](#lyc) != [LY](#ly) | [LYC](#lyc) = [LY](#ly) |
| 1-0 | Mode Flag                 | See below                | See below               |

| Mode bit 1 | Mode bit 0 | Description  | Clocks  |
|------------|------------|--------------|---------|
| 0          | 0          | H-Blank      | 201-207 |
| 0          | 1          | V-Blank      | 4560    |
| 1          | 0          | OAM Search   | 77-83   |
| 1          | 1          | LCD Transfer | 169-175 |

~~~
OAM Search + LCD Transfer + H-Blank = 456 clocks
144 * (OAM Search + LCD Transfer + H-Blank) + V-Blank = 70224 clocks
~~~

### SCY

* **Purpose**: Scroll Y
* **R/W**: Read & Write
* **Address**: `ff42`

Background viewport Y-offset.

### SCX

* **Purpose**: Scroll X
* **R/W**: Read & Write
* **Address**: `ff43`

Background viewport X-offset.

### LY

* **Purpose**: LCDC Y-coordinate
* **R/W**: Read-only
* **Address**: `ff44`

| Value   | Description                                          |
|---------|------------------------------------------------------|
| 0-143   | Index of line currently being transferred to the LCD |
| 144-153 | V-Blank, no line being transferred                   |

### LYC

* **Purpose**: LY Compare
* **R/W**: Read & Write
* **Address**: `ff45`

See [STAT](#stat).

### DMA

* **Purpose**: DMA Transfer to [OAM](#oam)
* **R/W**: Write-only
* **Address**: `ff46`

| Value   | Description                                        |
|---------|----------------------------------------------------|
| `00-f1` | Start DMA transfer from `xx00-xx9f` to [OAM](#oam) |
| `f2-ff` | Not supported                                      |

Only [HRAM](#hram) can be accessed during DMA. Therefore, DMA must be
launched from code in [HRAM](#hram)

DMA takes 160 microseconds. Therefore, execution must remain in
[HRAM](#hram) for 160 microseconds

~~~asm
; Somewhere not in HRAM

 ld a,$c0
 call $ff4c   ; start DMA transfer from $c000 - $c09f to OAM

; ...

; Code assumed to have been pre-copied into HRAM (starting at $ff4c)

 ld ($ff46),a ; start DMA transfer, a=start address/$100
 ld a,$28     ; delay...
wait:         ; total 5x40 cycles, approx 200ms
 dec a        ; 1 cycle
 jr nz,wait   ; 4 cycles
 ret
~~~

### BGP

* **Purpose**: BG Palette Colors
* **R/W**: Read & Write
* **Address**: `ff47`

| Bits | Purpose                     |
|------|-----------------------------|
| 7-6  | Color of palette value `11` |
| 5-4  | Color of palette value `10` |
| 3-2  | Color of palette value `01` |
| 1-0  | Color of palette value `00` |

| Color | Description |
|-------|-------------|
| `00`  | White       |
| `01`  | Light gray  |
| `10`  | Dark gray   |
| `11`  | Black       |

### OBP0

* **Purpose**: Object Palette #0 Colors
* **R/W**: Read & Write
* **Address**: `ff48`

See [BGP](#bgp).

Palette value `00` is used for transparency. Therefore, bits 1-0 are
unused.

### OBP1

* **Purpose**: Object Palette #1 Colors
* **R/W**: Read & Write
* **Address**: `ff49`

See [OBP0](#obp0).

Palette value `00` is used for transparency. Therefore, bits 1-0 are
unused.
