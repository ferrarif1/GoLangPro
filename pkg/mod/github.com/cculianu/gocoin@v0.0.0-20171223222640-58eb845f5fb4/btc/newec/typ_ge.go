package newec

import (
	"fmt"
)


type ge_t struct {
	x, y fe_t
	infinity bool
}

func (ge *ge_t) print(lab string) {
	if ge.infinity {
		fmt.Println(lab + " - infinity")
		return
	}
	fmt.Println(lab + ".x:", ge.x.String())
	fmt.Println(lab + ".y:", ge.y.String())
}

func (elem *ge_t) pubkey_parse(pub []byte) bool {
	if len(pub) == 33 && (pub[0] == 0x02 || pub[0] == 0x03) {
		elem.x.set_b32(pub[1:33])
		elem.set_xo(&elem.x, pub[0]==0x03)
	} else if len(pub) == 65 && (pub[0] == 0x04 || pub[0] == 0x06 || pub[0] == 0x07) {
		elem.x.set_b32(pub[1:33])
		elem.y.set_b32(pub[33:65])
		if (pub[0] == 0x06 || pub[0] == 0x07) && elem.y.is_odd() != (pub[0] == 0x07) {
			return false
		}
	} else {
		return false
	}
	return elem.is_valid()
}


func (r *ge_t) set_xy(x, y *fe_t) {
	r.infinity = false
	r.x = *x
	r.y = *y
}


func (a *ge_t) is_valid() bool {
	if a.infinity {
		return false
	}
	var y2, x3, c fe_t
	a.y.sqr(&y2)
	a.x.sqr(&x3); x3.mul(&x3, &a.x)
	c.set_int(7)
	x3.set_add(&c)
	y2.normalize()
	x3.normalize()
	return y2.equal(&x3)
}


func (r *ge_t) set_gej(a *gej_t) {
	var z2, z3 fe_t;
	a.z.inv_var(&a.z)
	a.z.sqr(&z2)
	a.z.mul(&z3, &z2)
	a.x.mul(&a.x, &z2)
	a.y.mul(&a.y, &z3)
	a.z.set_int(1)
	r.infinity = a.infinity
	r.x = a.x
	r.y = a.y
}

func (a *ge_t) precomp(w int) (pre []ge_t) {
	pre = make([]ge_t, (1 << (uint(w)-2)))
	pre[0] = *a;
	var x, d, tmp gej_t
	x.set_ge(a)
	x.double(&d)
	for i:=1 ; i<len(pre); i++ {
		d.add_ge(&tmp, &pre[i-1])
		pre[i].set_gej(&tmp)
	}
	return
}

func (a *ge_t) neg(r *ge_t) {
	r.infinity = a.infinity
	r.x = a.x
	r.y = a.y
	r.y.normalize()
	r.y.negate(&r.y, 1)
}


func (r *ge_t) set_xo(x *fe_t, odd bool) {
	var c, x2, x3 fe_t
	r.x = *x
	x.sqr(&x2)
	x.mul(&x3, &x2)
	r.infinity = false
	c.set_int(7)
	c.set_add(&x3)
	c.sqrt(&r.y)
	r.y.normalize()
	if r.y.is_odd() != odd {
		r.y.negate(&r.y, 1)
	}
}

/*
void secp256k1_ge_set_xo(secp256k1_ge_t *r, const secp256k1_fe_t *x, int odd) {
    r->x = *x
    secp256k1_fe_t x2; secp256k1_fe_sqr(&x2, x)
    secp256k1_fe_t x3; secp256k1_fe_mul(&x3, x, &x2)
    r->infinity = 0
    secp256k1_fe_t c; secp256k1_fe_set_int(&c, 7)
    secp256k1_fe_add(&c, &x3)
    secp256k1_fe_sqrt(&r->y, &c)
    secp256k1_fe_normalize(&r->y)
    if (secp256k1_fe_is_odd(&r->y) != odd)
        secp256k1_fe_negate(&r->y, &r->y, 1)
}
*/