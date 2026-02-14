# Blue Monochromatic Theme Preview

## Color Palette
- **Picton Blue**: `#3ACBE8` (brightest - lightest blue)
- **Battery Charged Blue**: `#1CA3DE` (bright)
- **Blue Cola**: `#0D85D8` (medium)
- **True Blue**: `#0160C9` (darker)
- **Crayola's Absolute Zero**: `#0041C7` (darkest)

---

## Theme 1: Blue Monochrome Dark (blue-monochrome-dark)
**Style**: Night mode with very dark background and bright accents

### Visual Layout Preview
```
┌─────────────────────────────────────────────────────────────┐
│ Chat Panel (BG: #0041C7 - darkest)                         │
│                                                               │
│  ┌─────────────────────────────┐                            │
│  │ User Message                │ (BG: #0D85D8 - medium)     │
│  │ Text: #3ACBE8 (brightest)   │                            │
│  └─────────────────────────────┘                            │
│                                                               │
│       ┌─────────────────────────────┐                       │
│       │ Agent Message               │ (BG: #0160C9 - darker)│
│       │ Text: #3ACBE8 (brightest)   │                       │
│       └─────────────────────────────┘                       │
│                                                               │
│  ┌─────────────────────────────────────────────────────────┐│
│  │ Input: Type here... (BG: #0041C7, Text: #3ACBE8)        ││
│  └─────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ Diagram Panel (BG: #0160C9 - slightly lighter)             │
│                                                               │
│  Diagram text: #3ACBE8 (brightest)                          │
│  Code highlights: #0D85D8 (medium)                          │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

### Color Mapping
| Property       | Color    | Hex Code | Usage                        |
|----------------|----------|----------|------------------------------|
| ChatBg         | Darkest  | #0041C7  | Main chat background         |
| DiagramBg      | Darker   | #0160C9  | Diagram panel background     |
| UserMsgBg      | Medium   | #0D85D8  | User message bubbles         |
| AgentMsgBg     | Darker   | #0160C9  | Agent message bubbles        |
| InputBg        | Darkest  | #0041C7  | Input field background       |
| TextPrimary    | Brightest| #3ACBE8  | Primary text (high contrast) |
| TextSecondary  | Bright   | #1CA3DE  | Secondary text               |
| TextMuted      | Medium   | #0D85D8  | Dimmed/muted text            |
| AccentUser     | Brightest| #3ACBE8  | User name/badges             |
| AccentAgent    | Bright   | #1CA3DE  | Agent name/badges            |
| AccentCode     | Medium   | #0D85D8  | Code/syntax highlights       |

**Characteristics**:
- Very dark background for comfortable night viewing
- High contrast between text and background
- Monochromatic blue palette throughout
- User messages pop with medium blue background
- Clear visual separation between panels

---

## Theme 2: Blue Monochrome (blue-monochrome)
**Style**: Balanced contrast with medium backgrounds and white text

### Visual Layout Preview
```
┌─────────────────────────────────────────────────────────────┐
│ Chat Panel (BG: #0160C9 - darker blue)                     │
│                                                               │
│  ┌─────────────────────────────┐                            │
│  │ User Message                │ (BG: #1CA3DE - bright)     │
│  │ Text: #FFFFFF (white)       │                            │
│  └─────────────────────────────┘                            │
│                                                               │
│       ┌─────────────────────────────┐                       │
│       │ Agent Message               │ (BG: #0D85D8 - medium)│
│       │ Text: #FFFFFF (white)       │                       │
│       └─────────────────────────────┘                       │
│                                                               │
│  ┌─────────────────────────────────────────────────────────┐│
│  │ Input: Type here... (BG: #0160C9, Text: #FFFFFF)        ││
│  └─────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────┐
│ Diagram Panel (BG: #0D85D8 - medium blue)                  │
│                                                               │
│  Diagram text: #FFFFFF (white)                              │
│  Code highlights: #3ACBE8 (brightest blue)                  │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

### Color Mapping
| Property       | Color         | Hex Code | Usage                        |
|----------------|---------------|----------|------------------------------|
| ChatBg         | Darker Blue   | #0160C9  | Main chat background         |
| DiagramBg      | Medium Blue   | #0D85D8  | Diagram panel background     |
| UserMsgBg      | Bright Blue   | #1CA3DE  | User message bubbles (pop!)  |
| AgentMsgBg     | Medium Blue   | #0D85D8  | Agent message bubbles        |
| InputBg        | Darker Blue   | #0160C9  | Input field background       |
| TextPrimary    | White         | #FFFFFF  | Primary text (max readability)|
| TextSecondary  | Brightest Blue| #3ACBE8  | Secondary text               |
| TextMuted      | Darkest Blue  | #0041C7  | Dimmed/muted text            |
| AccentUser     | Brightest Blue| #3ACBE8  | User name/badges             |
| AccentAgent    | Bright Blue   | #1CA3DE  | Agent name/badges            |
| AccentCode     | Medium Blue   | #0D85D8  | Code/syntax highlights       |

**Characteristics**:
- Medium blue backgrounds for balanced viewing
- White text for maximum readability
- Blue accents maintain monochromatic feel
- User messages have brightest blue background
- More traditional "app" look vs. dark theme

---

## Comparison

| Aspect              | Blue Monochrome Dark      | Blue Monochrome           |
|---------------------|---------------------------|---------------------------|
| **Vibe**            | Night mode, sleek         | Balanced, professional    |
| **Text Readability**| High (bright on dark)     | Maximum (white on blue)   |
| **Monochrome Purity**| 100% blue palette        | 90% (white text breaks it)|
| **Eye Strain**      | Low (dark background)     | Medium (brighter overall) |
| **Best For**        | Late night coding         | All-day use               |
| **Aesthetic**       | Cyberpunk, modern         | Professional, clean       |

---

## Implementation Notes

Both themes will be available in the theme selector as:
- **"Blue Monochrome Dark"** (`blue-monochrome-dark`)
- **"Blue Monochrome"** (`blue-monochrome`)

Users can switch between them using the `/theme` command and see live previews before selecting.
