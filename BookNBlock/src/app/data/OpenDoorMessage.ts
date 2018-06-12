export class OpenDoorMessage {
    doorId: number;
    // verschlüsselt (HyperLedger)
    renterId: string;
    renterPK: string;
    // verschlüsselt (HyperLedger)
    timestemp: string;
}
