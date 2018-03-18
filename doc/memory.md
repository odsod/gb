# Memory

## Memory map

### `0000-3fff` - ROM Bank 00

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

### `4000-7fff` - ROM Bank 01-FF

| Address     | Description   |
|-------------|---------------|
| `4000-7fff` | Cartridge ROM |

### `8000-9fff` - VRAM

| Address     | Description                        |
|-------------|------------------------------------|
| `8000-8fff` | Tile Pattern Table (unsigned mode) |
| `8800-97ff` | Tile Pattern Table (signed mode)   |
| `9800-9bff` | BG Tile Map 1                      |
| `9c00-9fff` | BG Tile Map 2                      |

### `a000-fdff` - RAM

| Address     |               |
|-------------|---------------|
| `a000-bfff` | Cartridge RAM |
| `c000-cfff` | Internal RAM  |
| `e000-fdff` | Echo RAM      |

### `fe00-feff` - OAM

| Address     | Description           |
|-------------|-----------------------|
| `fe00-fe9f` | 40 4-byte OAM entries |
| `fea0-feff` | Unused                |

### `ff00-ff4b` - I/O Registers

| Address     | Type            | Description                          |
|-------------|-----------------|--------------------------------------|
| `ff00`      | Joypad          | Joypad (P1) (R/W)                    |
| `ff01`      | Serial transfer | Serial transfer data (SB) (R/W)      |
| `ff02`      |                 | Serial transfer control (SC) (R/W)   |
| `ff03`      | Unused          |                                      |
| `ff04`      | Timer           | Divider Register (DIV) (R/W)         |
| `ff05`      |                 | Timer Counter (TIMA) (R/W)           |
| `ff06`      |                 | Timer Modulo (TMA) (R/W)             |
| `ff07`      |                 | Timer Control (TAC) (R/W)            |
| `ff08`      | Interrupts      | Interrupt Flag (IF) (R/W)            |
| `ff09`      | Unused          |                                      |
| `ff10-ff3f` | Sound           | TODO                                 |
| `ff40`      | Video           | [LCDC](#lcdc)                        |
| `ff41`      |                 | LCD Status (STAT) (R/W)              |
| `ff42`      |                 | Scroll Y (SCY) (R/W)                 |
| `ff43`      |                 | Scroll X (SCX) (R/W)                 |
| `ff44`      |                 | LCDC Y-Coordinate (LY) (R)           |
| `ff45`      |                 | LY Compare (LYC) (R/W)               |
| `ff46`      |                 | DMA Transfer Start Address (DMA) (W) |
| `ff47`      |                 | BG & Window Palette (BGP) (R/W)      |
| `ff48`      |                 | Object Palette 0 (OBP0) (R/W)        |
| `ff49`      |                 | Object Palette 1 (OBP1) (R/W)        |
| `ff4a`      |                 | Window Y Position (WY) (R/W)         |
| `ff4b`      |                 | Window X Position (WX) (R/W)         |

### `ff4c-fffe` - HRAM

| Address     | Description |
|-------------|-------------|
| `ff4c-fffe` | HRAM        |

### `ffff` - Interrupt Enable Register

| Address | Description                 |
|---------|-----------------------------|
| `ffff`  | Interrupt Enable (IE) (R/W) |

## I/O Registers

### LCDC

* **Name**: LCD Control
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
