export class OpenDoorMessage {
    doorId: number;
    // verschlüsselt (HyperLedger)
    renterPubkey: string;
    // verschlüsselt (HyperLedger)
    timestemp: string;
}
