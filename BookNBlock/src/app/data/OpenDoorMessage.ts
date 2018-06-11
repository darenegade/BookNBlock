export class OpenDoorMessage {
    doorId: number;
    // verschlüsselt (HyperLedger)
    renterId: string;
    renterPubkey: string;
    // verschlüsselt (HyperLedger)
    timestemp: string;
}
