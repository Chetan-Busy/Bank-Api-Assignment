const bodyParser = require("body-parser")
const express = require("express")
const cors = require("cors")
const BankRoutes = require("./node-routes/bank.route")
const BranchRoutes = require("./node-routes/branch.route")

const app = express()

app.use(cors())
app.use(bodyParser.json())

app.use("/bank" , BankRoutes)
app.use("/branch" ,BranchRoutes)


app.listen(8080 , () => {
    console.log("Nodejs Server running on port 8080")
})