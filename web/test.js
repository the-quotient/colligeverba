obj = {
    presses: [],
    push(ch) { 
        this.modal == -1 ? 
            ch.length == 1 ? 
            this.presses.push({ ch: ch, pos: this.presses.length }) 
            : null 
        : null 
    },
    pop() { this.modal == -1 ? this.presses.pop() : null },
    tr(pos) { 
        this.presses[pos].ch.includes('?') ? 
            this.presses[pos].ch += '?' 
            : this.presses[pos].ch = '?'; 
        this.presses[pos].ch.length > 2 ?
            this.presses[pos].ch = '?' 
            : null 
    },
    modal: -1,
    setModal(num) { this.modal = num },
    setOrs(pos, val) { 
        this.presses[pos] = { 
            ...this.presses[pos],
            ors: val
        };
        this.modal = -1 
    },
    getStr() {
        ret = '';
        this.presses.forEach((el) => {
            orsss = '' 
            if (el.ors !== undefined && el.ors !== null) {
                el.ors.split('').forEach( (orr) => orsss += orr+',' )
                orsss = orsss.substring(0, orsss.length-1)
            }
            thret = (el.ch.includes('?') && el.ch.length == 1) ? el.ch + '{' + orsss + '}' : el.ch;
            ret += thret;
        })
        return ret
    }
}
