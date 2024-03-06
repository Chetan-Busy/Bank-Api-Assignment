const { default: axios } = require("axios")
const express = require("express")

const router = express.Router()

const headers = {
    'Content-Type': 'application/json'
};

router.post("/" , async(req ,res) => {
    await axios.post("http://localhost:3000/branch", JSON.stringify(req.body), {headers}).then(
        response => {
            res.status(200).json({"message" : "Branch created successfully"})
        }
    ).catch(err => {
        res.status(400).json({"error" : err.response.data.error})
    })
})

router.delete("/:id" , async(req ,res) => {
    bankId = req.params.id
    await axios.delete(`http://localhost:3000/branch/${bankId}`, {headers}).then(
        response => {
            res.status(200).json({"message" : "Branch deleted successfully"})
        }
    ).catch(err => {
        res.status(400).json({"error" : err.response.data.error})
    })
})

router.get("/:id" , async(req ,res) => {
    bankId = req.params.id
    await axios.get(`http://localhost:3000/branch/${bankId}`, {headers}).then(
        response => {
            res.status(200).json({"response" : response.data})
        }
    ).catch(err => {
        res.status(400).json({"error" : err.response.data.error})
    })
})


router.patch("/" , async(req ,res) => {
    await axios.patch("http://localhost:3000/branch",JSON.stringify(req.body), {headers}).then(
        response => {
            res.status(200).json({"Message" : response.data})
        }
    ).catch(err => {
        res.status(400).json({"error" : err.response.data.error})
    })
})



module.exports = router