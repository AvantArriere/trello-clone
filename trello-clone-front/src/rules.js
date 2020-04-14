export default function () {
  return {
    email: [
      v => !!v || 'E-mail is required',
      v => /.+@.+\..+/.test(v) || 'E-mail must be valid',
    ],
    fullName: [
      v => !!v || 'Name is required.',
    ],
    password: [
      v => !!v || 'Password is required.',
      v => v.length > 8 || '8자 이상.',
      // 소문자, 대문자, 숫자, 특수문자가 한 개씩은 들어가야 한다.
      v => /(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])/.test(v) || '대문자, 소문자, 특수문자, 숫자를 포함.',
      v => !/(.*[^a-zA-Z0-9!@#$%^&*])/.test(v) || '대문자, 소문자, 특수문자, 숫자를 포함',
    ]
  }
}