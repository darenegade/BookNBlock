export class OpenDoorMessage {
    doorId: string;
    // verschlüsselt (HyperLedger)
    renterPubkey: string;
    // verschlüsselt (HyperLedger)
    timestamp: number;
    booking: number;
}
