import internal from "stream";

export interface BASKETInterface {
    ID: string,
    Amount: number;
    Add_time: Date | null;
    DOC_ID:number;
    DOCTOR:DOCTORInterface;
    MED_ID: number;
    MEDICINE: MEDICINEInterface;
    WHERE_ID: number;
    WHERE:WHEREInterface;
}
export interface DOCTORInterface {
    ID: string,
    Name: string,
}
export interface MEDICINEInterface {
    ID: number,
    Name: string,
    How: string,
    So: string,
    Unit: string,
}
export interface WHEREInterface {
    ID: number,
    Name: string,
}