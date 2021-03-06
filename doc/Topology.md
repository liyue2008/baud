#Baud Topology Definitions

//the global partition id range
message PartitionHashRange {
	uint16 ph_start = 2;
	uint16 ph_end = 3;
}

message ServerID {
	uint16 region
	uint16 cell
	uint16 rack
	uint16 node
}

enum ReplicaRole {
	UNKNOWN = 0;
	LEADER = 1;
	FOLLOWER = 2;
	BACKUP = 3
}

enum PartitionStatus {
	UNKNOWN = 0;
	SERVING = 1;
	SPLITTING = 2;
	MERGING = 3;
	DYING = 4;
}

message PartitionInfo {
	uint32 space = 1;
	PartitionHashRange id = 1;
	repeated ServerID replicas = 2;
	PartitionStatus status  = 3;
	PartitionHashRange parent = 4;
	repeated PartitionHashRange children = 5;
}

message ServerInfo {
	ServerID id = 1;
	string address = 2;
}

message AVZoneInfo {
	uint16 id = 1;
	string name = 2;
}

message RegionInfo {
	uint16 id = 1;
	string name = 2;
}

