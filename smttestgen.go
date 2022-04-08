package main

import (
    "encoding/binary"
    "os"

    "smttestgen/framework"
    "smttestgen/smtw"
)

func write(fileName string, tests []framework.Test, marshaller framework.Marshaller) {
    data, err := marshaller.Marshall(tests)
    if err != nil {
        panic("Unable to marshall tests!")
    }

    f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        panic("Unable to open file for writing!")
    }

    _, err = f.Write(data)
    if err != nil {
        panic("Unable to write data to file!")
    }
}

func populateTests(tests *[]framework.Test) {
    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Empty Root")
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 1")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 2")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x01"), []byte("DATA"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 3")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x01"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x02"), []byte("DATA"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 5")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x01"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x02"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x03"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x04"), []byte("DATA"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 10")
        for i := 0; i < 10; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 100")
        for i := 0; i < 100; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update With Repeated Inputs")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update Overwrite Key")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("CHANGE"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update Union")
        for i := 0; i < 5; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 10; i < 15; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 20; i < 25; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update Sparse Union")
        for i := 0; i < 5; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i*2))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update With Empty Data")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte(""))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update With Empty Data Performs Delete")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte(""))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 1 Delete 1")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Delete([]byte("\x00\x00\x00\x00"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 2 Delete 1")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x01"), []byte("DATA"))
        _, _ = smt.Delete([]byte("\x00\x00\x00\x01"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Update 10 Delete 5")
        for i := 0; i < 10; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 5; i < 10; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            _, _ = smt.Delete(bs)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Delete Non-existent Key")
        _, _ = smt.Update([]byte("\x00\x00\x00\x00"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x01"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x02"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x03"), []byte("DATA"))
        _, _ = smt.Update([]byte("\x00\x00\x00\x04"), []byte("DATA"))
        _, _ = smt.Delete([]byte("\x00\x00\x04\x00"))
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Interleaved Update Delete")
        for i := 0; i < 10; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 5; i < 15; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            _, _ = smt.Delete(bs)
        }
        for i := 10; i < 20; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 15; i < 25; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            _, _ = smt.Delete(bs)
        }
        for i := 20; i < 30; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 25; i < 35; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            _, _ = smt.Delete(bs)
        }
        *tests = append(*tests, smt.GetTest())
    }

    {
        smt := smtw.NewSparseMerkleTreeWrapper("Test Delete Sparse Union")
        for i := 0; i < 10; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i))
            d := []byte("DATA")
            _, _ = smt.Update(bs, d)
        }
        for i := 0; i < 5; i++ {
            bs := make([]byte, 4)
            binary.BigEndian.PutUint32(bs, uint32(i*2+1))
            _, _ = smt.Delete(bs)
        }
        *tests = append(*tests, smt.GetTest())
    }
}

func main() {
    var tests []framework.Test
    populateTests(&tests)

    jsonMarshaller := framework.NewJsonMarshaller()
    yamlMarshaller := framework.NewYamlMarshaller()
    write("./smt_test_spec.json", tests, jsonMarshaller)
    write("./smt_test_spec.yaml", tests, yamlMarshaller)
}
