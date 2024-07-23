### qr-generator
`A program for generating QR code for any request`


The program can generate a QR code based on the entered data, be it a word, a numerical value, or a link.
All saved codes are automatically saved in the `saved/"yyyy/mm/dd"`.
They can be distinguished by their Unique ID as well as the entered value, which is substituted into the name of the finished QR code.

---

#### Stack
- Go 1.22.5
- Qrcode 0.0
- uuid 1.6.0
- color 1.17
- bufio
- image
- regexp
- time

---

#### How to install and use it
1. Clone the remote repository and go to it:
```
git clone https://github.com/linkoffee/qr-generator.git
```
```
cd qr-generator
```
2. Install all dependencies in the main directory:
```
go mod tidy
```
3. Run the program, enter data to generate:
```
go run generate.go
```
```
Enter data to encode in QR code: your data here
```
4. If the program completed successfully, then your generated QR code will appear in the `saved/"yyyy-mm-dd"` directory:
```
QR code generation succeeded.
```
<img src="https://habrastorage.org/webt/h_/s0/y_/h_s0y_l7-xnf1a4cni_9rzgqljk.png" />
<img src="https://habrastorage.org/webt/6u/ag/sx/6uagsxuymqiv6m2b2uyt9fnqd7q.png" width="350"/>

---

Author: [linkoffee](https://github.com/linkoffee)
