// This file is automatically generated by qtc from "reg_valcode.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line library/email/tmpl/reg_valcode.qtpl:1
package tmpl

//line library/email/tmpl/reg_valcode.qtpl:1
import (
	"valerian/library/email/tmpl/layouts"
)

//line library/email/tmpl/reg_valcode.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line library/email/tmpl/reg_valcode.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line library/email/tmpl/reg_valcode.qtpl:7
type RegisterValcodeBody struct {
	Head     *layouts.EmailPageHead
	FromName string
	Valcode  string
}

//line library/email/tmpl/reg_valcode.qtpl:16
func (p *RegisterValcodeBody) StreamEmailHTML(qw422016 *qt422016.Writer) {
	//line library/email/tmpl/reg_valcode.qtpl:16
	qw422016.N().S(`
`)
	//line library/email/tmpl/reg_valcode.qtpl:17
	p.Head.StreamHeadHTML(qw422016)
	//line library/email/tmpl/reg_valcode.qtpl:17
	qw422016.N().S(`

<table role="presentation" border="0" cellpadding="0" cellspacing="0">
  <tr>
    <td>
      <p>你好</p>
      <p>欢迎使用他石笔记，你的注册验证码如下，请在 5 分钟内输入进行下一步操作。</p>

      <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="document-item">
        <tbody>
          <tr>
            <td colSpan="2" style="padding: 16px;">
              <p style="background-color:#fff; padding: 6px; font-size: 14px;">
              `)
	//line library/email/tmpl/reg_valcode.qtpl:30
	qw422016.E().S(p.Valcode)
	//line library/email/tmpl/reg_valcode.qtpl:30
	qw422016.N().S(`
              </p>
            </td>
          </tr>
        </tbody>
      </table>
    </td>
  </tr>
</table>
`)
	//line library/email/tmpl/reg_valcode.qtpl:39
	layouts.StreamFooterHTML(qw422016)
	//line library/email/tmpl/reg_valcode.qtpl:39
	qw422016.N().S(`

`)
//line library/email/tmpl/reg_valcode.qtpl:41
}

//line library/email/tmpl/reg_valcode.qtpl:41
func (p *RegisterValcodeBody) WriteEmailHTML(qq422016 qtio422016.Writer) {
	//line library/email/tmpl/reg_valcode.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line library/email/tmpl/reg_valcode.qtpl:41
	p.StreamEmailHTML(qw422016)
	//line library/email/tmpl/reg_valcode.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line library/email/tmpl/reg_valcode.qtpl:41
}

//line library/email/tmpl/reg_valcode.qtpl:41
func (p *RegisterValcodeBody) EmailHTML() string {
	//line library/email/tmpl/reg_valcode.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
	//line library/email/tmpl/reg_valcode.qtpl:41
	p.WriteEmailHTML(qb422016)
	//line library/email/tmpl/reg_valcode.qtpl:41
	qs422016 := string(qb422016.B)
	//line library/email/tmpl/reg_valcode.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
	//line library/email/tmpl/reg_valcode.qtpl:41
	return qs422016
//line library/email/tmpl/reg_valcode.qtpl:41
}
