pragma solidity 0.4.24;

contract applyContract {
    struct addrInfo {
        string finger;
        bool signed;
        bool used;
    }

    mapping (address => addrInfo) addrMap;
    address[] applicants;

    function apply(address addr, string finger) external returns(bool) {
        applicants.push(addr);
        addrMap[addr].used = true;

        return true;
    }

    function sign(address addr) external returns(bool)  {
        addrMap[addr].signed = true;

        return true;
    }

    function getApplicants() external view returns(address[]) {
        return applicants;
    }

    function getSigned() external view returns(address[]) {
        uint256 count;
        address[] memory temp = new address[](applicants.length);
        for (uint256 i = 0; i < applicants.length; i++) {
            address addr = applicants[i];
            if (addrMap[addr].signed) {
                temp[count] = addr;
                count++;
            }
        }

        address[] memory result = new address[](count);
        for (i = 0; i < count; i++) {
            result[i] = temp[i];
        }

        return result;
    }

    function getUnSigned() external view returns(address[]) {
        uint256 count;
        address[] memory temp = new address[](applicants.length);
        for (uint256 i = 0; i < applicants.length; i++) {
            address addr = applicants[i];
            if (addrMap[addr].used && !addrMap[addr].signed) {
                temp[count] = addr;
                count++;
            }
        }

        address[] memory result = new address[](count);
        for (i = 0; i < count; i++) {
            result[i] = temp[i];
        }

        return result;
    }
}
