import scripts from './listedCompany.json'

interface Script {
  id: number;
  companyName: string;
  symbol: string;
  securityName: string;
  status: string;
  companyEmail: string;
  website: string;
  sectorName: string;
  regulatoryBody: string;
  instrumentType: string

}

const myScripts = ["OMPL", "EBLD91", "CREST", "NMIC", "GSY"]

const FilterScriptDetail = (own: string[]) => {
  console.log("length of josn file", scripts.length)
  const filterScripts = scripts.filter((item: Script) => own.includes(item.symbol))
  console.log(filterScripts)
}

FilterScriptDetail(myScripts)
