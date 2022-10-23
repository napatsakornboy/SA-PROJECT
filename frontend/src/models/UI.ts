import internal from "stream";

export interface BASKETInterface {
    ID: string,
    Amount: number;
    Add_time: Date | null;
    DOCTOR_ID:number;
    DOCTOR:DOCTORInterface;
    MEDICINE_ID: number;
    MEDICINE: MEDICINEInterface;
    WHERE_ID: number;
    WHERE:WHEREInterface;
    Symtomp_ID: number;
    Symtomp:SymtompInterface;
}
export interface DOCTORInterface {
    ID: string,
    Name: string,
    Title: string,
}
export interface MEDICINEInterface {
    ID: string,
    Name: string,
    How: string,
    So: string,
    Unit: string,
}
export interface WHEREInterface {
    ID: string,
    Name: string,
}

export interface SymtompInterface {
    ID: string,
    Check_Date:  string,
	Temperature: number,
	Pressure:    number,
	Heartrate:   number,
	Comment:     string,
	MAPB_ID:     string,
	Check_Owner: string,
	Level_ID:    string,
	Medicine:    string,
    
    
}