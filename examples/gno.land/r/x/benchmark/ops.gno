package benchmark

/*
func OpDecl()

OpExec, type T (const-type gno.land/r/x/benchmark.T)
OpEval, (const-type gno.land/r/x/benchmark.T)
OpTypeDecl, type T (const-type gno.land/r/x/benchmark.T)
OpExec, bodyStmt[0/0/1]=var i<VPBlock(1,1)> (const-type int)
OpExec, var i<VPBlock(1,1)> (const-type int)
OpEval, (const-type int)
OpValueDecl, var i<VPBlock(1,1)> (const-type int)
OpExec, bodyStmt[0/0/2]=const c<VPBlock(1,2)> (const-type int) = (const (1 int))
OpExec, const c<VPBlock(1,2)> (const-type int) = (const (1 int))
OpEval, (const (1 int))
OpEval, (const-type int)
OpValueDecl, const c<VPBlock(1,2)> (const-type int) = (const (1 int))
OpExec, bodyStmt[0/0/3]=(end)
OpExec, return
OpReturnFromBlock, func()()
OpHalt
*/
func OpDecl() {
	type T int
	var i int
	const c int = 1
}

/*
func OpEvalInt()

OpEval, (const (1234567891 int))
OpDefine, i<VPBlock(1,0)> := (const (1234567891 int))
OpExec, bodyStmt[0/0/1]=i2<VPBlock(1,1)> := (const (1234567892 int))
OpEval, (const (1234567892 int))
OpDefine, i2<VPBlock(1,1)> := (const (1234567892 int))
OpExec, bodyStmt[0/0/2]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:OpEvalInt RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/
func OpEvalInt() {
	i := 1234567891
	i2 := 123_456_7892
}

/*
func OpEvalFloat()

OpEval, (const (1.23456789 float64))
OpDefine, f<VPBlock(1,0)> := (const (1.23456789 float64))
OpExec, bodyStmt[0/0/1]=f2<VPBlock(1,1)> := (const (1.23456789 float64))
OpEval, (const (1.23456789 float64))
OpDefine, f2<VPBlock(1,1)> := (const (1.23456789 float64))
OpExec, bodyStmt[0/0/2]=f3<VPBlock(1,2)> := (const (123.456789 float64))
OpEval, (const (123.456789 float64))
OpDefine, f3<VPBlock(1,2)> := (const (123.456789 float64))
OpExec, bodyStmt[0/0/3]=f4<VPBlock(1,3)> := (const (0.0123456789 float64))
OpEval, (const (0.0123456789 float64))
OpDefine, f4<VPBlock(1,3)> := (const (0.0123456789 float64))
OpExec, bodyStmt[0/0/4]=f5<VPBlock(1,4)> := (const (2048 float64))
OpEval, (const (2048 float64))
OpDefine, f5<VPBlock(1,4)> := (const (2048 float64))
OpExec, bodyStmt[0/0/5]=f6<VPBlock(1,5)> := (const (15.5 float64))
OpEval, (const (15.5 float64))
OpDefine, f6<VPBlock(1,5)> := (const (15.5 float64))
OpExec, bodyStmt[0/0/6]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:OpEvalFloat RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/
func OpEvalFloat() {
	// decimal
	f := 1.23456789
	f2 := 1.234_56789
	// exp
	f3 := 1.23456789e2
	f4 := 1.23456789e-2
	// hex
	f5 := 0x2.p10  // == 2048.0
	f6 := 0x1.Fp+0 // == 1.9375
}

/*
func StmtOps()

OpEval, (const (1000000 int))
OpDefine, i<VPBlock(1,0)> := (const (1000000 int))
OpExec, bodyStmt[0/0/1]=i<VPBlock(1,0)>++
OpInc, i<VPBlock(1,0)>++
OpExec, bodyStmt[0/0/2]=i<VPBlock(1,0)>--
OpDec, i<VPBlock(1,0)>--
OpExec, bodyStmt[0/0/3]=i<VPBlock(1,0)> += (const (1000000 int))
OpEval, (const (1000000 int))
OpAddAssgin, i<VPBlock(1,0)> += (const (1000000 int))
OpExec, bodyStmt[0/0/4]=i<VPBlock(1,0)> -= (const (1000000 int))
OpEval, (const (1000000 int))
OpSubAssgin, i<VPBlock(1,0)> -= (const (1000000 int))
OpExec, bodyStmt[0/0/5]=i<VPBlock(1,0)> *= (const (1 int))
OpEval, (const (1 int))
OpMulAssgin, i<VPBlock(1,0)> *= (const (1 int))
OpExec, bodyStmt[0/0/6]=i<VPBlock(1,0)> /= (const (1000000 int))
OpEval, (const (1000000 int))
OpQuoAssgin, i<VPBlock(1,0)> /= (const (1000000 int))
OpExec, bodyStmt[0/0/7]=i<VPBlock(1,0)> %= (const (3 int))
OpEval, (const (3 int))
OpRemAssgin, i<VPBlock(1,0)> %= (const (3 int))
OpExec, bodyStmt[0/0/8]=i<VPBlock(1,0)> &= (const (2 int))
OpEval, (const (2 int))
OpAddAssgin, i<VPBlock(1,0)> &= (const (2 int))
OpExec, bodyStmt[0/0/9]=i<VPBlock(1,0)> |= (const (12 int))
OpEval, (const (12 int))
OpBorAssgin, i<VPBlock(1,0)> |= (const (12 int))
OpExec, bodyStmt[0/0/10]=i<VPBlock(1,0)> &^= (const (6 int))
OpEval, (const (6 int))
OpBandnAssgin, i<VPBlock(1,0)> &^= (const (6 int))
OpExec, bodyStmt[0/0/11]=i<VPBlock(1,0)> ^= (const (14 int))
OpEval, (const (14 int))
OpXorAssgin, i<VPBlock(1,0)> ^= (const (14 int))
OpExec, bodyStmt[0/0/12]=i<VPBlock(1,0)> <<= (const (2 uint))
OpEval, (const (2 uint))
OpShlAssgin, i<VPBlock(1,0)> <<= (const (2 uint))
OpExec, bodyStmt[0/0/13]=i<VPBlock(1,0)> >>= (const (2 uint))
OpEval, (const (2 uint))
OpShrAssgin, i<VPBlock(1,0)> >>= (const (2 uint))
OpExec, bodyStmt[0/0/14]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:StmtOps RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/
func StmtOps() {
	i := 1_000_000
	i++            // 1,000,001
	i--            // 1,000,00
	i += 1_000_000 // 2_000_000
	i -= 1_000_000 // 1_000_000
	i *= 1         // 1,000,000
	i /= 1_000_000 // 1
	i %= 3         // 1
	i &= 2         // 01 & 10 = 0
	i |= 12        // 12( 1100 )
	i &^= 6        // 12&^6=8 (1100 &^ 0110 = 1000)
	i ^= 14        // 8^14 = 6 (1000 ^ 1110 = 0110)
	i <<= 2        // 24
	i >>= 2        // 6
}

func StmtOps1K() {
	i := 1_000_000
	for j := 0; j < 1000; j++ {
		i++            // 1,000,001
		i--            // 1,000,00
		i += 1_000_000 // 2_000_000
		i -= 1_000_000 // 1_000_000
		i *= 1         // 1,000,000
		i /= 1_000_000 // 1
		i %= 3         // 1
		i &= 2         // 01 & 10 = 0
		i |= 12        // 12( 1100 )
		i &^= 6        // 12&^6=8 (1100 &^ 0110 = 1000)
		i ^= 14        // 8^14 = 6 (1000 ^ 1110 = 0110)
		i <<= 2        // 24
		i >>= 2        // 6
	}
}

/*
func OpCall()

OpEval, Echo<VPBlock(3,8)>((const ("hello world!" string)))
OpEval, Echo<VPBlock(3,8)>
OpEval, (const ("hello world!" string))
OpPreCall, Echo<VPBlock(3,8)>((const ("hello world!" string)))
OpCall, Echo
OpExec, bodyStmt[0/0/-2]=(init)
OpEval, s<VPBlock(1,0)>
OpReturn, [FRAME FUNC:Echo RECV:(undefined) (1 args) 4/1/0/2/2 LASTPKG:gno.land/r/x/benchmark LASTRLM:Realm{Path:"gno.land/r/x/benchmark",Time:3}#707D4A13D8A59C3A9220761016E2B0AF5FFCBC5A]
OpDefine, a<VPBlock(1,0)> := Echo<VPBlock(3,8)>((const ("hello world!" string)))
OpExec, bodyStmt[0/0/1]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:OpCall RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/

func OpCall() {
	a := Echo("hello world!")
}

func Echo(s string) string {
	return s
}

/*
func ControlOps()

OpExec, var a<VPBlock(1,0)> (const-type int)
OpEval, (const-type int)
OpValueDecl, var a<VPBlock(1,0)> (const-type int)
OpExec, bodyStmt[0/0/1]=var b<VPBlock(1,1)> (const-type int)
OpExec, var b<VPBlock(1,1)> (const-type int)
OpEval, (const-type int)
OpValueDecl, var b<VPBlock(1,1)> (const-type int)
OpExec, bodyStmt[0/0/2]=if (const-type bool)(a<VPBlock(2,0)> > (const (0 int))) { b<VPBlock(2,1)>++ } else { b<VPBlock(2,1)>-- }
OpEval, (const-type bool)(a<VPBlock(2,0)> > (const (0 int)))
OpEval, (const-type bool)
OpEval, a<VPBlock(2,0)> > (const (0 int))
OpEval, a<VPBlock(2,0)>
OpEval, (const (0 int))
OpGtr, (0 int) | (0 int) | false
OpPreCall, (const-type bool)(a<VPBlock(2,0)> > (const (0 int)))
OpConvert, Value: (false <untyped> bool) | Type: bool
OpIfCond, if (const-type bool)(a<VPBlock(2,0)> > (const (0 int))) { b<VPBlock(2,1)>++ } else { b<VPBlock(2,1)>-- }
OpExec, bodyStmt[0/0/-2]=(init)
OpDec, b<VPBlock(2,1)>--
OpExec, bodyStmt[0/0/1]=(end)
OpPopBlock, Block(ID:0000000000000000000000000000000000000000:0,Addr:0x1400ad725a0,Source:if (const-type bool)(a<VPBlock(2...,Parent:0x1400ad723c0)
OpExec, bodyStmt[0/0/3]=switch b<VPBlock(2,1)> { case (const (0 int)): a<VPBlock(2,0)>++; case (const (1 int)): a<VPBlock(2,0)>--; default: a<VPBlock(2,0)> = (const (2 int)) }
OpEval, b<VPBlock(2,1)>
OpSwitchClause, switch b<VPBlock(2,1)> { case (const (0 int)): a<VPBlock(2,0)>++; case (const (1 int)): a<VPBlock(2,0)>--; default: a<VPBlock(2,0)> = (const (2 int)) }
OpEval, (const (0 int))
OpSwitchClauseCase, (0 int) | (-1 int)
OpSwitchClause, switch b<VPBlock(2,1)> { case (const (0 int)): a<VPBlock(2,0)>++; case (const (1 int)): a<VPBlock(2,0)>--; default: a<VPBlock(2,0)> = (const (2 int)) }
OpEval, (const (1 int))
OpSwitchClauseCase, (1 int) | (-1 int)
OpSwitchClause, switch b<VPBlock(2,1)> { case (const (0 int)): a<VPBlock(2,0)>++; case (const (1 int)): a<VPBlock(2,0)>--; default: a<VPBlock(2,0)> = (const (2 int)) }
OpExec, bodyStmt[0/0/-2]=(init)
OpEval, (const (2 int))
OpAssgin, a<VPBlock(2,0)> = (const (2 int))
OpExec, bodyStmt[0/0/1]=(end)
OpPopBlock, Block(ID:0000000000000000000000000000000000000000:0,Addr:0x1400ad72780,Source:switch b<VPBlock(2,1)> { case (c...,Parent:0x1400ad723c0)
OpPopFrameAndReset, [FRAME LABEL:  3/1/0/2/2]
OpExec, bodyStmt[0/0/4]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:ControlOps RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/
func ControlOps() {
	var a int
	var b int
	if a > 0 {
		b++
	} else {
		b--
	}
	switch b {

	case 0:
		a++
	case 1:
		a--
	default:
		a = 2
	}
}

/*
func OpDefer()

OpExec, var b<VPBlock(1,1)> (const-type int)
OpEval, (const-type int)
OpValueDecl, var b<VPBlock(1,1)> (const-type int)
OpExec, bodyStmt[0/0/1]=defer func func(){ b<VPBlock(2,1)>++ }()
OpEval, func func(){ b<VPBlock(2,1)>++ }
OpEval, func()
OpFuncType, func()
OpFuncLit, func func(){ b<VPBlock(2,1)>++ }
OpDefer, defer func func(){ b<VPBlock(2,1)>++ }()
OpExec, bodyStmt[0/0/2]=return b<VPBlock(1,1)>
OpEval, b<VPBlock(1,1)>
OpReturnToBlock, [FRAME FUNC:OpDefer RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpReturnCallDefers
OpExec, bodyStmt[0/0/-2]=(init)
OpInc, b<VPBlock(2,1)>++
OpExec, bodyStmt[0/0/1]=(end)
OpReturnCallDefers
OpReturnFromBlock, [FRAME FUNC:OpDefer RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/
func OpDefer() int {
	var b int

	defer func() {
		b++
	}()
	return b
}
/*
func OpUnary()


*/
func OpUnary(){

	a := 1
	b := -a
	b = +a
	b = ^a

	c := true
	d := !c

}
/*
func OpBinary()

OpEval, (const (1000000 int))
OpDefine, a<VPBlock(1,0)> := (const (1000000 int))
OpExec, bodyStmt[0/0/1]=b<VPBlock(1,1)> := (const (1000001 int))
OpEval, (const (1000001 int))
OpDefine, b<VPBlock(1,1)> := (const (1000001 int))
OpExec, bodyStmt[0/0/2]=var c<VPBlock(1,2)> (const-type bool)
OpExec, var c<VPBlock(1,2)> (const-type bool)
OpEval, (const-type bool)
OpValueDecl, var c<VPBlock(1,2)> (const-type bool)
OpExec, bodyStmt[0/0/3]=var d<VPBlock(1,3)> (const-type int)
OpExec, var d<VPBlock(1,3)> (const-type int)
OpEval, (const-type int)
OpValueDecl, var d<VPBlock(1,3)> (const-type int)
OpExec, bodyStmt[0/0/4]=c<VPBlock(1,2)> = (const (true bool))
OpEval, (const (true bool))
OpAssgin, c<VPBlock(1,2)> = (const (true bool))
OpExec, bodyStmt[0/0/5]=c<VPBlock(1,2)> = (const (false bool))
OpEval, (const (false bool))
OpAssgin, c<VPBlock(1,2)> = (const (false bool))
OpExec, bodyStmt[0/0/6]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> == b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> == b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> == b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpEql, (1000000 int) == (1000001 int) is false
OpPreCall, (const-type bool)(a<VPBlock(1,0)> == b<VPBlock(1,1)>)
OpConvert, Value: (false <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> == b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/7]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> != b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> != b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> != b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpNeq, (1000000 int) != (1000001 int) is true
OpPreCall, (const-type bool)(a<VPBlock(1,0)> != b<VPBlock(1,1)>)
OpConvert, Value: (true <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> != b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/8]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> < b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> < b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> < b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpLss, (1000000 int) < (1000001 int) is true
OpPreCall, (const-type bool)(a<VPBlock(1,0)> < b<VPBlock(1,1)>)
OpConvert, Value: (true <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> < b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/9]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> <= b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> <= b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> <= b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpLeq, (1000000 int) <= (1000001 int) is true
OpPreCall, (const-type bool)(a<VPBlock(1,0)> <= b<VPBlock(1,1)>)
OpConvert, Value: (true <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> <= b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/10]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> > b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> > b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> > b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpGtr, (1000000 int) > (1000001 int) is false
OpPreCall, (const-type bool)(a<VPBlock(1,0)> > b<VPBlock(1,1)>)
OpConvert, Value: (false <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> > b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/11]=c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> >= b<VPBlock(1,1)>)
OpEval, (const-type bool)(a<VPBlock(1,0)> >= b<VPBlock(1,1)>)
OpEval, (const-type bool)
OpEval, a<VPBlock(1,0)> >= b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpGeq, (1000000 int) >= (1000001 int) is false
OpPreCall, (const-type bool)(a<VPBlock(1,0)> >= b<VPBlock(1,1)>)
OpConvert, Value: (false <untyped> bool) | Type: bool
OpAssgin, c<VPBlock(1,2)> = (const-type bool)(a<VPBlock(1,0)> >= b<VPBlock(1,1)>)
OpExec, bodyStmt[0/0/12]=d<VPBlock(1,3)> = a<VPBlock(1,0)> + b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> + b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpAdd, (1000000 int) + (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> + b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/13]=d<VPBlock(1,3)> = a<VPBlock(1,0)> - b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> - b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpSub, (1000000 int) - (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> - b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/14]=d<VPBlock(1,3)> = a<VPBlock(1,0)> | b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> | b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpBor, (1000000 int) | (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> | b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/15]=d<VPBlock(1,3)> = a<VPBlock(1,0)> ^ b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> ^ b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpXor, (1000000 int) ^ (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> ^ b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/16]=d<VPBlock(1,3)> = a<VPBlock(1,0)> * b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> * b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpMul, (1000000 int) * (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> * b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/17]=d<VPBlock(1,3)> = a<VPBlock(1,0)> / b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> / b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpQuo, (1000000 int) / (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> / b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/18]=d<VPBlock(1,3)> = a<VPBlock(1,0)> % b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> % b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpRem, (1000000 int) % (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> % b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/19]=d<VPBlock(1,3)> = (const (63 int)) << (const-type uint)(a<VPBlock(1,0)>)
OpEval, (const (63 int)) << (const-type uint)(a<VPBlock(1,0)>)
OpEval, (const (63 int))
OpEval, (const-type uint)(a<VPBlock(1,0)>)
OpEval, (const-type uint)
OpEval, a<VPBlock(1,0)>
OpPreCall, (const-type uint)(a<VPBlock(1,0)>)
OpConvert, Value: (1000000 int) | Type: uint
OpShl, (63 int) << (1000000 uint)
OpAssgin, d<VPBlock(1,3)> = (const (63 int)) << (const-type uint)(a<VPBlock(1,0)>)
OpExec, bodyStmt[0/0/20]=d<VPBlock(1,3)> = a<VPBlock(1,0)> >> (const (63 uint))
OpEval, a<VPBlock(1,0)> >> (const (63 uint))
OpEval, a<VPBlock(1,0)>
OpEval, (const (63 uint))
OpShr, (1000000 int) >> (63 uint)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> >> (const (63 uint))
OpExec, bodyStmt[0/0/21]=d<VPBlock(1,3)> = a<VPBlock(1,0)> & b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> & b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpBand, (1000000 int) & (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> & b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/22]=d<VPBlock(1,3)> = a<VPBlock(1,0)> &^ b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)> &^ b<VPBlock(1,1)>
OpEval, a<VPBlock(1,0)>
OpEval, b<VPBlock(1,1)>
OpBandn, (1000000 int) &^ (1000001 int)
OpAssgin, d<VPBlock(1,3)> = a<VPBlock(1,0)> &^ b<VPBlock(1,1)>
OpExec, bodyStmt[0/0/23]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:OpBinary RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt


*/
func OpBinary() {
	a := 1_000_000
	b := 1_000_001

	var c bool
	var d int
	// boolean
	c = true || false
	c = true && false
	c = a == b
  c = a != b
	c = a < b
	c = a <= b
	c = a > b
	c = a >= b
	//
	d = a + b  // 1
	d = a - b  // 0
	d = a | b  // 1
	d = a ^ b  // 1
	d = a * b  // 1,000,001,000,000
	d = a / b
	d = a % b
	d = 63 << a
	d = a >> 63
	d = a & b
	d = a &^ b

}

type foo struct{
    i int
}
func (f foo) bark() {

}

type  dog interface{

      bark()
}

type foofighter struct{
	f foo
}
/* func ExprOps()
OpEval, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpEval, [(const (2 int))](const-type int)
OpEval, (const-type int)
OpEval, (const (2 int))
OpArrayType, [(const (2 int))](const-type int)
OpCompositeLit, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpEval, (const (0 int))
OpEval, (const (1 int))
OpArrayLit, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpDefine, a<VPBlock(1,0)> := [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpExec, bodyStmt[0/0/1]=a2<VPBlock(1,1)> := [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpEval, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpEval, [(const (2 int))](const-type int)
OpEval, (const-type int)
OpEval, (const (2 int))
OpArrayType, [(const (2 int))](const-type int)
OpCompositeLit, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpEval, (const (0 int))
OpEval, (const (1 int))
OpArrayLit, [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpDefine, a2<VPBlock(1,1)> := [(const (2 int))](const-type int){(const (0 int)), (const (1 int))}
OpExec, bodyStmt[0/0/2]=m<VPBlock(1,2)> := (const (make func(t type{},z ...interface{})( map[int]int)))(map[(const-type int)] (const-type int))
OpEval, (const (make func(t type{},z ...interface{})( map[int]int)))(map[(const-type int)] (const-type int))
OpEval, (const (make func(t type{},z ...interface{})( map[int]int)))
OpEval, map[(const-type int)] (const-type int)
OpEval, (const-type int)
OpEval, (const-type int)
OpMapType, (typeval{int} type{})
OpPreCall, (const (make func(t type{},z ...interface{})( map[int]int)))(map[(const-type int)] (const-type int))
OpCall, make
OpCallNativeBody, make
OpReturn, [FRAME FUNC:make RECV:(undefined) (1 args) 4/1/0/2/2 LASTPKG:gno.land/r/x/benchmark LASTRLM:Realm{Path:"gno.land/r/x/benchmark",Time:3}#707D4A13D8A59C3A9220761016E2B0AF5FFCBC5A]
OpDefine, m<VPBlock(1,2)> := (const (make func(t type{},z ...interface{})( map[int]int)))(map[(const-type int)] (const-type int))
OpExec, bodyStmt[0/0/3]=s<VPBlock(1,3)> := [](const-type int){(const (0 int)), (const (1 int)), (const (2 int)), (const (3 int)), (const (4 int)), (const (5 int)), (const (6 int)), (const (7 int)), (const (8 int)), (const (9 int))}
OpEval, [](const-type int){(const (0 int)), (const (1 int)), (const (2 int)), (const (3 int)), (const (4 int)), (const (5 int)), (const (6 int)), (const (7 int)), (const (8 int)), (const (9 int))}
OpEval, [](const-type int)
OpEval, (const-type int)
OpSliceType, [](const-type int)
OpCompositeLit, [](const-type int){(const (0 int)), (const (1 int)), (const (2 int)), (const (3 int)), (const (4 int)), (const (5 int)), (const (6 int)), (const (7 int)), (const (8 int)), (const (9 int))}
OpEval, (const (0 int))
OpEval, (const (1 int))
OpEval, (const (2 int))
OpEval, (const (3 int))
OpEval, (const (4 int))
OpEval, (const (5 int))
OpEval, (const (6 int))
OpEval, (const (7 int))
OpEval, (const (8 int))
OpEval, (const (9 int))
OpSliceLit, [](const-type int){(const (0 int)), (const (1 int)), (const (2 int)), (const (3 int)), (const (4 int)), (const (5 int)), (const (6 int)), (const (7 int)), (const (8 int)), (const (9 int))}
OpDefine, s<VPBlock(1,3)> := [](const-type int){(const (0 int)), (const (1 int)), (const (2 int)), (const (3 int)), (const (4 int)), (const (5 int)), (const (6 int)), (const (7 int)), (const (8 int)), (const (9 int))}
OpExec, bodyStmt[0/0/4]=s2<VPBlock(1,4)> := [](const-type int){(const (9 int)): (const (90 int))}
OpEval, [](const-type int){(const (9 int)): (const (90 int))}
OpEval, [](const-type int)
OpEval, (const-type int)
OpSliceType, [](const-type int)
OpCompositeLit, [](const-type int){(const (9 int)): (const (90 int))}
OpEval, (const (9 int))
OpEval, (const (90 int))
OpSliceLit2, [](const-type int){(const (9 int)): (const (90 int))}
OpDefine, s2<VPBlock(1,4)> := [](const-type int){(const (9 int)): (const (90 int))}
OpExec, bodyStmt[0/0/5]=f<VPBlock(1,5)> := foo<VPBlock(3,0)>{i<VPField(0,0,i)>: (const (1 int))}
OpEval, foo<VPBlock(3,0)>{i<VPField(0,0,i)>: (const (1 int))}
OpEval, foo<VPBlock(3,0)>
OpCompositeLit, foo<VPBlock(3,0)>{i<VPField(0,0,i)>: (const (1 int))}
OpEval, (const (1 int))
OpStructLit, foo<VPBlock(3,0)>{i<VPField(0,0,i)>: (const (1 int))}
OpDefine, f<VPBlock(1,5)> := foo<VPBlock(3,0)>{i<VPField(0,0,i)>: (const (1 int))}
OpExec, bodyStmt[0/0/6]=ff<VPBlock(1,6)> := foofighter<VPBlock(3,2)>{f<VPField(0,0,f)>: f<VPBlock(1,5)>}
OpEval, foofighter<VPBlock(3,2)>{f<VPField(0,0,f)>: f<VPBlock(1,5)>}
OpEval, foofighter<VPBlock(3,2)>
OpCompositeLit, foofighter<VPBlock(3,2)>{f<VPField(0,0,f)>: f<VPBlock(1,5)>}
OpEval, f<VPBlock(1,5)>
OpStructLit, foofighter<VPBlock(3,2)>{f<VPField(0,0,f)>: f<VPBlock(1,5)>}
OpDefine, ff<VPBlock(1,6)> := foofighter<VPBlock(3,2)>{f<VPField(0,0,f)>: f<VPBlock(1,5)>}
OpExec, bodyStmt[0/0/7]=b<VPBlock(1,7)> := a<VPBlock(1,0)>[(const (0 int))]
OpEval, a<VPBlock(1,0)>[(const (0 int))]
OpEval, a<VPBlock(1,0)>
OpEval, (const (0 int))
OpIndex1, (array[(0 int),(1 int)] [2]int)
OpDefine, b<VPBlock(1,7)> := a<VPBlock(1,0)>[(const (0 int))]
OpExec, bodyStmt[0/0/8]=b<VPBlock(1,7)>, _<VPBlock(0,0)> = m<VPBlock(1,2)>[(const (0 int))]
OpEval, m<VPBlock(1,2)>[(const (0 int))]
OpEval, m<VPBlock(1,2)>
OpEval, (const (0 int))
OpIndex2, (map{} map[int]int)
OpAssgin, b<VPBlock(1,7)>, _<VPBlock(0,0)> = m<VPBlock(1,2)>[(const (0 int))]
OpExec, bodyStmt[0/0/9]=b<VPBlock(1,7)> = f<VPBlock(1,5)>.i
OpEval, f<VPBlock(1,5)>.i
OpEval, f<VPBlock(1,5)>
OpSelector, f<VPBlock(1,5)>.i
OpAssgin, b<VPBlock(1,7)> = f<VPBlock(1,5)>.i
OpExec, bodyStmt[0/0/10]=subs<VPBlock(1,8)> := s<VPBlock(1,3)>[(const (1 int)):(const (5 int)):(const (10 int))]
OpEval, s<VPBlock(1,3)>[(const (1 int)):(const (5 int)):(const (10 int))]
OpEval, s<VPBlock(1,3)>
OpEval, (const (1 int))
OpEval, (const (5 int))
OpEval, (const (10 int))
OpSlice, s<VPBlock(1,3)>[(const (1 int)):(const (5 int)):(const (10 int))]
OpDefine, subs<VPBlock(1,8)> := s<VPBlock(1,3)>[(const (1 int)):(const (5 int)):(const (10 int))]
OpExec, bodyStmt[0/0/11]=ptr<VPBlock(1,9)> := &(a2<VPBlock(1,1)>[(const (0 int))])
OpEval, &(a2<VPBlock(1,1)>[(const (0 int))])
OpEval, a2<VPBlock(1,1)>
OpEval, (const (0 int))
OpRef, &(a2<VPBlock(1,1)>[(const (0 int))])
OpDefine, ptr<VPBlock(1,9)> := &(a2<VPBlock(1,1)>[(const (0 int))])
OpExec, bodyStmt[0/0/12]=b<VPBlock(1,7)> = *(ptr<VPBlock(1,9)>)
OpEval, *(ptr<VPBlock(1,9)>)
OpEval, ptr<VPBlock(1,9)>
OpStar, (&0x1400dfd25f0.(*int) *int)
OpAssgin, b<VPBlock(1,7)> = *(ptr<VPBlock(1,9)>)
OpExec, bodyStmt[0/0/13]=var d<VPBlock(1,10)> dog<VPBlock(3,1)>
OpExec, var d<VPBlock(1,10)> dog<VPBlock(3,1)>
OpEval, dog<VPBlock(3,1)>
OpValueDecl, var d<VPBlock(1,10)> dog<VPBlock(3,1)>
OpExec, bodyStmt[0/0/14]=d<VPBlock(1,10)> = f<VPBlock(1,5)>
OpEval, f<VPBlock(1,5)>
OpAssgin, d<VPBlock(1,10)> = f<VPBlock(1,5)>
OpExec, bodyStmt[0/0/15]=f<VPBlock(1,5)> = d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpEval, d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpEval, d<VPBlock(1,10)>
OpEval, foo<VPBlock(3,0)>
OpTypeAssert1, concrete type (struct{(1 int)} gno.land/r/x/benchmark.foo)
OpAssgin, f<VPBlock(1,5)> = d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpExec, bodyStmt[0/0/16]=f<VPBlock(1,5)>, ok<VPBlock(1,11)> := d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpEval, d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpEval, d<VPBlock(1,10)>
OpEval, foo<VPBlock(3,0)>
OpTypeAssert2, concrete type (struct{(1 int)} gno.land/r/x/benchmark.foo)
OpDefine, f<VPBlock(1,5)>, ok<VPBlock(1,11)> := d<VPBlock(1,10)>.(foo<VPBlock(3,0)>)
OpExec, bodyStmt[0/0/17]=(end)
OpExec, return
OpReturnFromBlock, [FRAME FUNC:ExprOps RECV:(undefined) (0 args) 1/0/0/0/1 LASTPKG:main LASTRLM:Realm(nil)]
OpHalt
*/

func ExprOps() {
	  a := [2]int{0,1} // OpArrayLit

		a2 := [...]int{0,1} // same as [2]int
		m := make(map[int]int) // OpMapLit
		s := []int{0,1,2,3,4,5,6,7,8,9} // OpSliceLit

		s2 := []int{9:90} // OpSliceLit2
		f :=foo {i: 1} // OpStructLit
		ff :=foofighter{f: f} // OpCompositeLit

	 b := a[0] // OpIndex1
	 b, _ = m[0] // OpIndex2
	 b = f.i // OpSelector

	 subs := s[1:5:10] // OpSlice

	 ptr := &a2[0] //OpRef
   b = *ptr  // OpStar

	 var d dog
 	 d = f
   f = d.(foo) // OpTypeAssert1 concrete type
//	 d = f.(d) // OpTypeAssert1 interface
	 f, ok := d.(foo) //OpTypeAssert2

}
