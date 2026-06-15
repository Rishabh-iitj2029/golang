const walletMoney = 100
const toPay = 120

const paymentStatus = async () => {
  try {

    if ((walletMoney-toPay) >= 0) {
      const res = await fetch("http://localhost:4000/api/v1/paymentstatus",{
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json'
              },
              body: JSON.stringify({
                  amountPaid: toPay,
                  balance:walletMoney-toPay,
                  statuscode:200
              })
    })
    console.log(await res.text())
    } else {
      const res = await fetch("http://localhost:4000/api/v1/paymentstatus",{
              method: 'POST',
              headers: {
                  'Content-Type': 'application/json'
              },
              body: JSON.stringify({
                  amountPaid: toPay,
                  balance:walletMoney-toPay,
                  statuscode:402
              })
    })
    console.log(await res.text())
    }
  
  } catch (error) {
    console.log(error)
  }
}

paymentStatus();